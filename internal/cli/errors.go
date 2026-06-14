package cli

import (
	"errors"
	"flag"
	"fmt"
)

var errKubeconfigNotFound = errors.New("kube config not found")
var errUnknownCommand = errors.New("unknown command")

// This function is called if unknown verbs are used: allowed verbs are get, describe, watch, check
func printUsageVerbs(flags *flag.FlagSet) {
	fmt.Fprintf(flags.Output(), "Usage: kubewatch <command> [flags]\n\n")
	fmt.Fprintf(flags.Output(), "Commands:\n")
	fmt.Fprintf(flags.Output(), "  get       Get Kubernetes resources\n")
	fmt.Fprintf(flags.Output(), "  describe  Describe Kubernetes resources\n")
	fmt.Fprintf(flags.Output(), "  watch     Watch Kubernetes resources\n")
	fmt.Fprintf(flags.Output(), "  check     Run a one-time health check\n\n")
	fmt.Fprintf(flags.Output(), "Flags:\n")
	flags.PrintDefaults()
}

// This function is called if an unknown resource type is used.
// Allowed resource types are pods, deployments, services, nodes, namespaces,
// daemonsets, statefulsets, jobs, cronjobs, configmaps, secrets, ingresses,
// persistentvolumes and persistentvolumeclaims.
func printUsageResourceTypes() {
	fmt.Println("Supported Resource Types:")
	fmt.Println("  pod, pods")
	fmt.Println("  deployment, deploy, deployments")
	fmt.Println("  service, svc, services")
	fmt.Println("  node, nodes")
	fmt.Println("  namespace, ns, namespaces")
	fmt.Println("  daemonset, ds, daemonsets")
	fmt.Println("  statefulset, sts, statefulsets")
	fmt.Println("  configmap, cm, configmaps")
	fmt.Println("  secret, secrets")
	fmt.Println("  ingress, ing, ingresses")
	fmt.Println("  pv, persistentvolumes")
	fmt.Println("  pvc, persistentvolumeclaims")
}
