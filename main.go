package main

import (
	"fmt"

	"github.com/karlma/mtr/cli"
)

func main() {
	err := cli.RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
