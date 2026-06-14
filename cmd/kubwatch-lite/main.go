package main

import (
	"fmt"
	"os"
)

func main() {
	kubeconfig, err := run(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(kubeconfig)

}
