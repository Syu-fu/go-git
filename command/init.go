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

	if err := os.Mkdir(gitPath, 0755); err != nil {
		fmt.Println(err)
		os.Exit(129)
	}

	if err := os.Mkdir(path.Join(gitPath, "objects"), 0755); err != nil {
		fmt.Println(err)
		os.Exit(129)
	}

	if err := os.MkdirAll(path.Join(gitPath, "refs", "tags"), 0755); err != nil {
		fmt.Println(err)
		os.Exit(129)
	}
	if err := os.MkdirAll(path.Join(gitPath, "refs", "heads"), 0755); err != nil {
		fmt.Println(err)
		os.Exit(129)
	}

	des, err := os.Create(path.Join(gitPath, "description"))
	if err != nil {
		panic("createErr:")
	}
	if _, err := des.WriteString("Unnamed repository; edit this file 'description' to name the repository."); err != nil {
		fmt.Println(err)
		os.Exit(129)
	}

	head, err := os.Create(path.Join(gitPath, "HEAD"))
	if err != nil {
		panic("createErr:")
	}
	if _, err := head.WriteString("ref: refs/heads/\n"); err != nil {
		fmt.Println(err)
		os.Exit(129)
	}
	if err := os.MkdirAll(path.Join(gitPath, "refs", "heads"), 0755); err != nil {
		fmt.Println(err)
		os.Exit(129)
	}

	if _, err := os.Create(path.Join(gitPath, "config")); err != nil {
		fmt.Println(err)
		os.Exit(129)
	}
	cfg, err := ini.Load(path.Join(gitPath, "config"))
	if err != nil {
		fmt.Println(err)
		os.Exit(129)
	}
	cfg.Section("core").Key("repositoryformatversion").SetValue("0")
	cfg.Section("core").Key("filemode").SetValue("true")
	cfg.Section("core").Key("bare").SetValue("false")
	cfg.Section("core").Key("logallrefupdates").SetValue("true")
	if err := cfg.SaveToIndent(path.Join(gitPath, "config"), "	"); err != nil {
		fmt.Println(err)
		os.Exit(129)
	}

	fmt.Println("Initialized empty Git repository in " + absPath + "/")
}
