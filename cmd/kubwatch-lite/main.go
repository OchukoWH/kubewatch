package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var errUnknownCommand = errors.New("unknown command")
var errKubeconfigNotFound = errors.New("kube config not found")

func main() {
	kubeconfig, err := run(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(kubeconfig)

}

// Function to get the kubeconfig, this fuction is overriden if a --kubeconfig is provided.
func defaultKubeconfig() (string, error) {
	if kubeconfig := os.Getenv("KUBECONFIG"); kubeconfig != "" {
		return kubeconfig, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(home, ".kube", "config")

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("%w: %s", errKubeconfigNotFound, path)
		}
		return "", err
	}

	return path, nil
}

// Main entry into the bin
func run(args []string) (string, error) {
	kubeconfig, err := defaultKubeconfig()
	var namespace string
	if err != nil {
		return "", err
	}

	flags := flag.NewFlagSet("kubewatch", flag.ContinueOnError)
	flags.StringVar(&namespace, "namespace", "default", "namespace to inspect")
	flags.StringVar(&namespace, "n", "default", "namespace to inspect")

	verb, remainingArgs, ok := splitVerb(args)
	if !ok {
		printUsage(flags)
		return "", fmt.Errorf("missing command")
	}

	if err := flags.Parse(remainingArgs); err != nil {
		return "", err
	}

	if err := runVerb(verb, namespace, flags.Args()); err != nil {
		if errors.Is(err, errUnknownCommand) {
			printUsage(flags)
		}

		return "", err
	}

	return kubeconfig, nil
}

// Checks if the arg is a flag e.g --namespace
func isFlag(arg string) bool {
	return strings.HasPrefix(arg, "-")
}

// Checks specifically if the flag is namespace
func isNamespaceFlag(arg string) bool {
	return arg == "-n" || arg == "--namespace"
}

// This function splits the cmd arguments and parses the verb e.g get, describe, etc
func splitVerb(args []string) (string, []string, bool) {
	for index := 0; index < len(args); index++ {
		arg := args[index]
		// if arg == "-n" || arg == "--namespace" {
		// 	index++
		// 	continue
		// }
		isNameSpace := isNamespaceFlag(arg)
		if isNameSpace {
			index++
			continue
		}

		// if strings.HasPrefix(arg, "-") {
		// 	continue
		// }
		isSubcommand := isFlag(arg)
		if isSubcommand {
			continue
		}

		if arg != "" {
			remainingArgs := append(args[:index:index], args[index+1:]...)
			return arg, remainingArgs, true
		}
	}

	return "", args, false
}

// func splitResourceType(args []string) (string, []string, bool) {
// 	for index := 0; index < len(args); index++ {

// 	}
// }

// Uses the correct verb to get the required resources
func runVerb(verb string, namespace string, args []string) error {
	switch verb {
	case "get":
		return runGet(namespace, args)
	case "describe":
		return runDescribe(namespace, args)
	case "watch":
		return runWatch(namespace, args)
	case "check":
		return runCheck(namespace, args)
	default:
		return fmt.Errorf("%w %q", errUnknownCommand, verb)
	}
}

// Function for the get verb
func runGet(namespace string, args []string) error {
	fmt.Printf("get: namespace=%s args=%v\n", namespace, args)
	return nil
}

// Function for the describe
func runDescribe(namespace string, args []string) error {
	fmt.Printf("describe: namespace=%s args=%v\n", namespace, args)
	return nil
}

// Function for the watch verb
func runWatch(namespace string, args []string) error {
	fmt.Printf("watch: namespace=%s args=%v\n", namespace, args)
	return nil
}

// Function for the check verb
func runCheck(namespace string, args []string) error {
	fmt.Printf("check: namespace=%s args=%v\n", namespace, args)
	return nil
}

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
