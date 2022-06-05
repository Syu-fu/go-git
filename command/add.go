package command

import (
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func Add(filenameList []string) {
	for _, v := range filenameList {
		createGitObjectBlob(v)
	}
}

func getObjectId(store string) string {
	h := sha1.New()
	_, err := io.WriteString(h, store)
	if err != nil {
		fmt.Println(err)
		os.Exit(129)
	}
	return hex.EncodeToString(h.Sum(nil))
}

func createGitObjectBlob(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		panic("Error in fileopen" + err.Error())
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(129)
	}
	content := string(b)
	header := "blob " + fmt.Sprint(len(content)) + "\x00"
	addFilepath := getObjectId(header + content)

	gitPath := gitRootPath()
	if err := os.Mkdir(path.Join(gitPath, "objects", addFilepath[:2]), 0755); err != nil {
		fmt.Println(err)
		os.Exit(129)
	}

	nf, err := os.Create(path.Join(gitPath, "objects", addFilepath[:2], addFilepath[2:]))
	if err != nil {
		panic("createErr:")
	}
	nw := zlib.NewWriter(nf)
	defer nw.Close()
	if _, err := nw.Write([]byte(header + content)); err != nil {
		fmt.Println(err)
		os.Exit(129)
	}
}
