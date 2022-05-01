package command

import (
	"compress/zlib"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type GitObject struct {
	Type     string
	FileSize string
	Content  string
}

func CatFile(hash string, option string) {
	gitObject := DecodeGitObject(hash)
	if option == "type" {
		fmt.Println(gitObject.Type)
	} else if option == "size" {
		fmt.Println(gitObject.FileSize)
	} else if option == "pretty-print" {
		fmt.Print(gitObject.Content)
	} else {
		panic("option error")
	}
}

func DecodeGitObject(hash string) GitObject {
	path := HashToFilePath(hash)
	object := decompress(path)

	return ParseGitObject(object)
}

func decompress(filepath string) string {
	f, err := os.Open(filepath)
	if err != nil {
		panic("Error in fileopen" + err.Error())
	}
	defer f.Close()

	z, err := zlib.NewReader(f)
	if err != nil {
		panic("Error in zlib:" + err.Error())
	}
	defer z.Close()

	b := make([]byte, 2048)
	z.Read(b)
	return string(b)
}

func HashToFilePath(hash string) string {
	dir := hash[0:2]
	file := hash[2:]
	return string(filepath.Join(GitRootPath(), "objects", dir, file))
}

func GitRootPath() string {
	path := RepoRootPath(".")
	return filepath.Join(path, ".git")
}

func RepoRootPath(currentPath string) string {
	if _, err := os.Stat(filepath.Join(currentPath, ".git")); err == nil {
		return currentPath
	}

	parentPath := filepath.Join(currentPath, "..")
	if parentPath == currentPath {
		panic("No git repo directory")
	}

	return RepoRootPath(parentPath)
}

func ParseGitObject(object string) GitObject {
	spaceIndex := strings.Index(object, " ")
	objectType := object[0:spaceIndex]

	nullIndex := strings.Index(object, "\x00")
	objectSize := object[spaceIndex+1 : nullIndex]

	objectSizeInt, err := strconv.Atoi(objectSize)
	if err != nil {
		panic("Object size parse Error")
	}
	content := object[nullIndex+1 : nullIndex+1+objectSizeInt]

	var gitObject GitObject
	switch objectType {
	case "blob":
		gitObject.Type = objectType
		gitObject.FileSize = objectSize
		gitObject.Content = content
		return gitObject
	case "commit":
		gitObject.Type = objectType
		gitObject.FileSize = objectSize
		gitObject.Content = content
		return gitObject
	case "tree":
		gitObject.Type = objectType
		gitObject.FileSize = objectSize
		gitObject.Content = content
		return gitObject
	default:
		panic("No git file")
	}
}
