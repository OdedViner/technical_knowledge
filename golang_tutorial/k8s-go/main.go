package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/oviner/ClusterPath/auth/kubeconfig", "location to kubeconfig file.")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		//handel error
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		//handel error
	}
	ctx := context.Background()
	pods, err := clientset.CoreV1().Pods("openshift-storage").List(ctx, metav1.ListOptions{})
	if err != nil {
		//handel error
	}
	fmt.Println("pod in openshift-storage ns.")
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}

	deployments, err := clientset.AppsV1().Deployments("openshift-storage").List(ctx, metav1.ListOptions{})
	if err != nil {
		//handel error
	}
	fmt.Println("deployment in openshift-storage ns.")
	for _, d := range deployments.Items {
		fmt.Println(d.Name)
	}
}
