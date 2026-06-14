package cli

import (
	"fmt"

	"github.com/OchukoWH/kubewatch/internal/kube"
)

func runGet(ctx *CLIContext) error {
	clientset, err := kube.NewClient(ctx.Kubeconfig)
	if err != nil {
		return err
	}

	pods, err := kube.GetPods(clientset, ctx.Namespace)
	if err != nil {
		return err
	}

	for _, pod := range pods {
		fmt.Printf(
			"%s\t%s\t%s\n",
			pod.Name,
			pod.Status.Phase,
			pod.Spec.NodeName,
		)
	}

	return nil
}
