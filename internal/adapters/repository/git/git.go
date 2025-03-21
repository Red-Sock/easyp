package git

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"go.redsock.ru/protopack/internal/adapters/repository"
)

var _ repository.Repo = (*gitRepo)(nil)

// gitRepo implements repository.Repo interface
type gitRepo struct {
	// remoteURL full repository remoteURL address with schema
	remoteURL string
	// cacheDir local cache directory for store repository
	cacheDir string
	// console for call external commands
	console Console
}

const (
	// for omitted package version. HEAD is git key word.
	gitLatestVersionRef = "HEAD"
	// tag prefix on output of ls-remote command
	gitRefsTagPrefix = "refs/tags/"
)

// Some links from go mod:
// cmd/go/internal/modfetch/codehost/git.go:65 - create work dir
// cmd/go/internal/modfetch/codehost/git.go:137 - git's struct

// Console temporary interface for console commands, must be replaced from core.Console.
type Console interface {
	RunCmd(ctx context.Context, dir string, command string, commandParams ...string) (string, error)
}

// New returns gitRepo instance
// remote: full remoteURL address without schema
func New(ctx context.Context, remote string, cacheDir string, console Console) (repository.Repo, error) {
	r := &gitRepo{
		remoteURL: getRemote(remote),
		cacheDir:  cacheDir,
		console:   console,
	}

	if _, err := os.Stat(filepath.Join(r.cacheDir, "objects")); err == nil {
		// repo is already exists
		return r, nil
	}

	if _, err := r.console.RunCmd(ctx, r.cacheDir, "git", "init", "--bare"); err != nil {
		return nil, fmt.Errorf("adapters.RunCmd (init): %w", err)
	}

	_, err := r.console.RunCmd(ctx, r.cacheDir, "git", "remote", "add", "origin", r.remoteURL)
	if err != nil {
		return nil, fmt.Errorf("adapters.RunCmd (add origin): %w", err)
	}

	return r, nil
}

func getRemote(name string) string {
	return "https://" + name
}
