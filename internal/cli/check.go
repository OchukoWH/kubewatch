package cli

import "fmt"

// Function for the check verb
func runCheck(namespace string, args []string) error {
	fmt.Printf("check: namespace=%s args=%v\n", namespace, args)
	return nil
}
