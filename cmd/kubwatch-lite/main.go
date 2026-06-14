package main

import (
	"fmt"
	"os"

	cli "github.com/OchukoWH/kubewatch/internal/cli"
)

func main() {
	kubeconfig, err := cli.Run(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(kubeconfig)

}
