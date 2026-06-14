package main

import (
	"os"

	cli "github.com/OchukoWH/kubewatch/internal/cli"
)

func main() {
	cli.Run(os.Args[1:])

}
