package kube

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetPods(clientset *kubernetes.Clientset, namespace string) ([]corev1.Pod, error) {
	podList, err := clientset.
		CoreV1().
		Pods(namespace).
		List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return podList.Items, nil
}
