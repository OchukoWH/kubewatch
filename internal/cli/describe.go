package cli

import "fmt"

// Function for the describe
func runDescribe(namespace string, args []string) error {
	fmt.Printf("describe: namespace=%s args=%v\n", namespace, args)
	return nil
}
