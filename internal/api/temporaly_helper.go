package api

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/samber/lo"

	"go.redsock.ru/protopack/internal/adapters/console"
	"go.redsock.ru/protopack/internal/adapters/go_git"
	lockfile "go.redsock.ru/protopack/internal/adapters/lock_file"
	moduleconfig "go.redsock.ru/protopack/internal/adapters/module_config"
	"go.redsock.ru/protopack/internal/adapters/storage"
	"go.redsock.ru/protopack/internal/config"
	"go.redsock.ru/protopack/internal/core"
	"go.redsock.ru/protopack/internal/rules"
)

var (
	ErrPathNotAbsolute = errors.New("path is not absolute")
)

const (
	envEasypPath     = "EASYPPATH"
	defaultEasypPath = ".easyp"
)

func errExit(code int, msg string, args ...any) {
	slog.Info(msg, args...)
	os.Exit(code)
}

// getEasypPath return path for cache, modules storage
func getEasypPath() (string, error) {
	easypPath := os.Getenv(envEasypPath)
	if easypPath == "" {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("os.UserHomeDir: %w", err)
		}
		easypPath = filepath.Join(userHomeDir, defaultEasypPath)
	}

	easypPath, err := filepath.Abs(easypPath)
	if err != nil {
		return "", ErrPathNotAbsolute
	}

	slog.Debug("Use storage", "path", easypPath)

	return easypPath, nil
}

func buildCore(_ context.Context, cfg config.Config, dirWalker core.DirWalker) (*core.Core, error) {
	lintRules, ignoreOnly, err := rules.New(cfg.Lint)
	if err != nil {
		return nil, fmt.Errorf("cfg.BuildLinterRules: %w", err)
	}

	lockFile := lockfile.New(dirWalker)
	easypPath, err := getEasypPath()
	if err != nil {
		return nil, fmt.Errorf("getEasypPath: %w", err)
	}

	store := storage.New(easypPath, lockFile)

	moduleCfg := moduleconfig.New()

	currentProjectGitWalker := go_git.New()

	breakingCheckConfig := core.BreakingCheckConfig{
		IgnoreDirs:    cfg.BreakingCheck.Ignore,
		AgainstGitRef: cfg.BreakingCheck.AgainstGitRef,
	}

	app := core.New(
		lintRules,
		cfg.Lint.Ignore,
		cfg.Deps,
		ignoreOnly,
		slog.Default(), // TODO: remove global state
		lo.Map(cfg.Generate.Plugins, func(p config.Plugin, _ int) core.Plugin {
			return core.Plugin{
				Name:    p.Name,
				Out:     p.Out,
				Options: p.Opts,
			}
		}),
		core.Inputs{
			Dirs: lo.Filter(lo.Map(cfg.Generate.Inputs, func(i config.Input, _ int) string {
				return i.Directory
			}), func(s string, _ int) bool {
				return s != ""
			}),
			InputGitRepos: lo.Filter(lo.Map(cfg.Generate.Inputs, func(i config.Input, _ int) core.InputGitRepo {
				return core.InputGitRepo{
					URL:          i.GitRepo.URL,
					SubDirectory: i.GitRepo.SubDirectory,
					Out:          i.GitRepo.Out,
				}
			}), func(i core.InputGitRepo, _ int) bool {
				return i.URL != ""
			}),
		},
		console.New(),
		store,
		moduleCfg,
		lockFile,
		currentProjectGitWalker,
		breakingCheckConfig,
		cfg.Generate.ProtoRoot,
		cfg.Generate.GenerateOutDirs,
	)

	return app, nil
}
