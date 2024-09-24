package base

import (
	"os"
	"path/filepath"
	"strings"
)

func Dir(target string) string {
	dest, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dest = filepath.ToSlash(dest)
	target = filepath.ToSlash(target)

	destDirs := strings.Split(dest, "/")
	targetDirs := strings.Split(target, "/")

	var p1, p2 int
	n1, n2 := len(destDirs), len(targetDirs)
	for p1 < n1 && p2 < n2 {
		if destDirs[p1] == targetDirs[p2] {
			targetDirs = targetDirs[1:]
			p2++
		}
		p1++
	}

	if len(targetDirs) > 0 {
		dest = filepath.Join(
			dest,
			filepath.Join(targetDirs...),
		)
	}

	return dest
}
