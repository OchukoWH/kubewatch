package cli

import (
	"errors"
	"flag"
	"fmt"
)

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
