package cli

import "fmt"

// Function for the get verb
func runGet(namespace string, args []string) error {
	fmt.Printf("get: namespace=%s args=%v\n", namespace, args)
	return nil
}
