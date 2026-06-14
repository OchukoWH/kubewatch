package cli

import "fmt"

// Function for the check verb
func runCheck(ctx *CLIContext) error {
	fmt.Printf("check: namespace=%s args=%v\n", ctx.Namespace, ctx.Verb)
	return nil
}
