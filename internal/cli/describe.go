package cli

import "fmt"

// Function for the describe
func runDescribe(ctx *CLIContext) error {
	fmt.Printf("describe: namespace=%s args=%v\n", ctx.Namespace, ctx.Verb)
	return nil
}
