package cli

import "fmt"

// Function for the get verb
func runGet(ctx *CLIContext) error {
	fmt.Printf("get: namespace=%s args=%v\n", ctx.Namespace, ctx.Verb)
	return nil
}
