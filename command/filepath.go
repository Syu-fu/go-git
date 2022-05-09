package command

import (
	"os"
	"path/filepath"
)

func gitRootPath() string {
	path := repoRootPath(".")
	return filepath.Join(path, ".git")
}

func repoRootPath(currentPath string) string {
	if _, err := os.Stat(filepath.Join(currentPath, ".git")); err == nil {
		return currentPath
	}

	parentPath := filepath.Join(currentPath, "..")
	if parentPath == currentPath {
		panic("No git repo directory")
	}

	return repoRootPath(parentPath)
}

func hashToFilePath(hash string) string {
	dir := hash[0:2]
	file := hash[2:]
	return string(filepath.Join(gitRootPath(), "objects", dir, file))
}
