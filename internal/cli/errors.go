package cli

import (
	"errors"
	"flag"
	"fmt"
)

var errKubeconfigNotFound = errors.New("kube config not found")
var errUnknownCommand = errors.New("unknown command")

// This function is called if unknown cmd flags are used
func printUsage(flags *flag.FlagSet) {
	fmt.Fprintf(flags.Output(), "Usage: kubewatch <command> [flags]\n\n")
	fmt.Fprintf(flags.Output(), "Commands:\n")
	fmt.Fprintf(flags.Output(), "  get       Get Kubernetes resources\n")
	fmt.Fprintf(flags.Output(), "  describe  Describe Kubernetes resources\n")
	fmt.Fprintf(flags.Output(), "  watch     Watch Kubernetes resources\n")
	fmt.Fprintf(flags.Output(), "  check     Run a one-time health check\n\n")
	fmt.Fprintf(flags.Output(), "Flags:\n")
	flags.PrintDefaults()
}
