package main

import (
	"flag"
	"fmt"
	"go-git/command"
	"os"
)

func main() {
	catFile := flag.NewFlagSet("cat-file", flag.ExitOnError)
	t := catFile.Bool("t", false, "show object type")
	s := catFile.Bool("s", false, "show object size")
	p := catFile.Bool("p", false, "pretty-print object's content")

	init := flag.NewFlagSet("init", flag.ExitOnError)
	b := init.String("b", "main", "Use the specified name for the initial branch in the newly created repository.")

	switch os.Args[1] {
	case "cat-file":
		catFile.Parse(os.Args[2:])

		option := "pretty-print"
		if flag.NFlag() > 1 {
			panic("too many options")
		}
		if *t {
			option = "type"
		} else if *s {
			option = "size"
		} else if *p {
			option = "pretty-print"
		}

		command.CatFile(catFile.Arg(0), option)
	case "init":
		init.Parse(os.Args[2:])

		command.Init(*b)
	case "add":
		command.Add(os.Args[2:])
	default:
		fmt.Println("go-git: '" + os.Args[1] + "' is not a git command. See 'go-git --help'.")
	}

}
