package cli

var ResourceAliases = map[string]string{
	"pod":  "pods",
	"pods": "pods",

	"deployment":  "deployments",
	"deploy":      "deployments",
	"deployments": "deployments",

	"service":  "services",
	"svc":      "services",
	"services": "services",

	"node":  "nodes",
	"nodes": "nodes",

	"namespace":  "namespaces",
	"ns":         "namespaces",
	"namespaces": "namespaces",

	"daemonset":  "daemonsets",
	"ds":         "daemonsets",
	"daemonsets": "daemonsets",

	"statefulset":  "statefulsets",
	"sts":          "statefulsets",
	"statefulsets": "statefulsets",

	"job":  "jobs",
	"jobs": "jobs",

	"cronjob":  "cronjobs",
	"cj":       "cronjobs",
	"cronjobs": "cronjobs",

	"configmap":  "configmaps",
	"cm":         "configmaps",
	"configmaps": "configmaps",

	"secret":  "secrets",
	"secrets": "secrets",

	"ingress":   "ingresses",
	"ing":       "ingresses",
	"ingresses": "ingresses",

	"pv":                "persistentvolumes",
	"persistentvolume":  "persistentvolumes",
	"persistentvolumes": "persistentvolumes",

	"pvc":                    "persistentvolumeclaims",
	"persistentvolumeclaim":  "persistentvolumeclaims",
	"persistentvolumeclaims": "persistentvolumeclaims",
}

type CLIContext struct {
	Verb       string
	Resource   string
	Namespace  string
	Kubeconfig string
	Args       []string
}
