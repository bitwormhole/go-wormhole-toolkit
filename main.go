package main

import (
	"fmt"

	"github.com/bitwormhole/go-wormhole-toolkit/tools/autoconfig"
)

func main() {
	fmt.Println("hello, autoconfig")

	args := &autoconfig.CommandArgs{}
	context, err := autoconfig.Run(args)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("autoconfig package dir:" + context.PackageDirectory.Path.Path())
		fmt.Println("Success.")
	}

	fmt.Println("done.")
}
