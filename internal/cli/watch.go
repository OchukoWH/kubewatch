package cli

import "fmt"

// Function for the watch verb
func runWatch(ctx *CLIContext) error {
	fmt.Printf("watch: namespace=%s args=%v\n", ctx.Namespace, ctx.Verb)
	return nil
}
