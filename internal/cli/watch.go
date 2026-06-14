package cli

import "fmt"

// Function for the watch verb
func runWatch(namespace string, args []string) error {
	fmt.Printf("watch: namespace=%s args=%v\n", namespace, args)
	return nil
}
