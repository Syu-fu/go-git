package command

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/ini.v1"
)

func Init(initialBranchName string) {

	gitPath := path.Join(".", ".git")
	absPath, _ := filepath.Abs(gitPath)
	if _, err := os.Stat(gitPath); err != nil {
		fmt.Println("Reinitialized existing Git repository in " + absPath)
		os.Exit(0)
	}

	os.Mkdir(gitPath, 0755)

	os.Mkdir(path.Join(gitPath, "objects"), 0755)
	os.MkdirAll(path.Join(gitPath, "refs", "tags"), 0755)
	os.MkdirAll(path.Join(gitPath, "refs", "heads"), 0755)

	des, err := os.Create(path.Join(gitPath, "description"))
	if err != nil {
		panic("createErr:")
	}
	des.WriteString("Unnamed repository; edit this file 'description' to name the repository.")

	head, err := os.Create(path.Join(gitPath, "HEAD"))
	if err != nil {
		panic("createErr:")
	}
	head.WriteString("ref: refs/heads/\n")
	os.MkdirAll(path.Join(gitPath, "refs", "heads"), 0755)

	os.Create(path.Join(gitPath, "config"))
	cfg, err := ini.Load(path.Join(gitPath, "config"))
	cfg.Section("core").Key("repositoryformatversion").SetValue("0")
	cfg.Section("core").Key("filemode").SetValue("true")
	cfg.Section("core").Key("bare").SetValue("false")
	cfg.Section("core").Key("logallrefupdates").SetValue("true")
	cfg.SaveToIndent(path.Join(gitPath, "config"), "	")

	fmt.Println("Initialized empty Git repository in " + absPath + "/")
}
