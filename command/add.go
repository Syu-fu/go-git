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

func getObjectId(store string) string {
	h := sha1.New()
	io.WriteString(h, store)
	return hex.EncodeToString(h.Sum(nil))
}

func createGitObjectBlob(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		panic("Error in fileopen" + err.Error())
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	content := string(b)
	header := "blob " + fmt.Sprint(len(content)) + "\x00"
	fmt.Println(getObjectId(header + content))
	addFilepath := getObjectId(header + content)

	gitPath := GitRootPath()
	os.Mkdir(path.Join(gitPath, "objects", addFilepath[:2]), 0755)

	nf, err := os.Create(path.Join(gitPath, "objects", addFilepath[:2], addFilepath[2:]))
	if err != nil {
		panic("createErr:")
	}
	nw := zlib.NewWriter(nf)
	defer nw.Close()
	nw.Write([]byte(header + content))
}
