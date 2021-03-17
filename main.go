package main

import (
	"log"

	"github.com/adarshnbidari/template/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	cmd.Execute()

	if err := doc.GenMarkdownTree(cmd.RootCmd, "./"); err != nil {

		log.Fatal(err)

	}

}
