package testdata

import (
	"github.com/golang/glog"
	"os"
	"path/filepath"
)

func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}

	v, err := goPackagePath("github.com/syunkitada/go-sample/pkg/simple_app/testdata/")
	if err != nil {
		glog.Fatalf("Error finding github.com/syunkitada/go-sample/pkg/simple_app/testdata/ directory: %v", err)
	}

	return filepath.Join(v, rel)
}

func goPackagePath(pkg string) (path string, err error) {
	gp := os.Getenv("GOPATH")
	if gp == "" {
		return path, os.ErrNotExist
	}

	for _, p := range filepath.SplitList(gp) {
		dir := filepath.Join(p, "src", filepath.FromSlash(pkg))
		fi, err := os.Stat(dir)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return "", err
		}
		if !fi.IsDir() {
			continue
		}
		return dir, nil
	}
	return path, os.ErrNotExist
}
