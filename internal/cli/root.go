package cli

import (
	"flag"
	"fmt"
)

// Uses the correct verb to get the required resources
func runVerb(ctx *CLIContext) error {
	switch ctx.Verb {
	case "get":
		return runGet(ctx)
	case "describe":
		return runDescribe(ctx)
	default:
		return fmt.Errorf("%w %q", errUnknownCommand, ctx.Verb)
	}
}

func Run(args []string) error {
	ctx, err := parseArgs(args)
	if err != nil {
		return err
	}

	return runVerb(ctx)
}

func parseArgs(args []string) (*CLIContext, error) {
	kubeconfig, err := defaultKubeconfig()
	if err != nil {
		return nil, err
	}

	var namespace string

	flags := flag.NewFlagSet("kubewatch", flag.ContinueOnError)
	flags.StringVar(&namespace, "namespace", "default", "namespace to inspect")
	flags.StringVar(&namespace, "n", "default", "namespace to inspect")

	verb, remainingArgs, ok := splitVerb(args)
	if !ok {
		printUsageVerbs(flags)
		return nil, fmt.Errorf("missing command")
	}

	if err := flags.Parse(remainingArgs); err != nil {
		return nil, err
	}

	resource, remainingArgs, ok := getResourceType(remainingArgs)
	if !ok {
		printUsageResourceTypes()
		return nil, fmt.Errorf("missing resource type")
	}

	return &CLIContext{
		Verb:       verb,
		Resource:   resource,
		Namespace:  namespace,
		Kubeconfig: kubeconfig,
		Args:       remainingArgs,
	}, nil
}
