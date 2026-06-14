package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Checks if the arg is a flag e.g --namespace
func isFlag(arg string) bool {
	return strings.HasPrefix(arg, "-")
}

// Checks specifically if the flag is namespace
func isNamespaceFlag(arg string) bool {
	return arg == "-n" || arg == "--namespace"
}

// Check if verb, we use this function across various other functions.
func isVerb(arg string) bool {
	return arg == "get" || arg == "check" || arg == "describe" || arg == "watch"
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

// Get the resource type we want like pod, deploy, etc
func getResourceType(args []string) (string, []string, bool) {
	for index := 0; index < len(args); index++ {
		arg := args[index]
		isNameSpace := isNamespaceFlag(arg)
		if isNameSpace {
			index++
			continue
		}

		isSubcommand := isFlag(arg)
		if isSubcommand {
			continue
		}

		verb := isVerb(arg)
		if verb {
			continue
		}

		if arg != "" {
			resource, ok := ResourceAliases[arg]
			if ok {
				remainingArgs := append(args[:index:index], args[index+1:]...)
				return resource, remainingArgs, true
			}
		}
	}

	return "", args, false
}
