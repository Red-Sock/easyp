package rules

import (
	"path/filepath"
	"strings"

	"github.com/easyp-tech/easyp/internal/lint"
)

var _ lint.Rule = (*PackageDirectoryMatch)(nil)

// PackageDirectoryMatch is a rule for checking consistency of directory and package names.
type PackageDirectoryMatch struct {
	Root string
}

// Validate implements core.Rule.
func (d *PackageDirectoryMatch) Validate(protoInfo lint.ProtoInfo) []error {
	var res []error

	preparePath := filepath.Dir(strings.TrimPrefix(protoInfo.Path, d.Root))
	expectedPackage := strings.Replace(preparePath, "/", ".", -1)

	for _, pkgInfo := range protoInfo.Info.ProtoBody.Packages {
		if pkgInfo.Name != expectedPackage {
			res = append(res, buildError(pkgInfo.Meta.Pos, protoInfo.Path, lint.ErrPackageIsNotMatchedWithPath))
		}
	}

	if len(res) == 0 {
		return nil
	}

	return res
}