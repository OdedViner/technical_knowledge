package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/oviner/ClusterPath/auth/kubeconfig", "location to kubeconfig file.")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		//handel error
		fmt.Printf("error %s building config from flags\n", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("error %s, getting inclusterconfig", err.Error())
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		//handel error
		fmt.Printf("error %s creating clientset\n", err.Error())
	}
	ctx := context.Background()
	pods, err := clientset.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		//handel error
		fmt.Printf("error %s, while listing all the pods from default namespace\n", err.Error())
	}
	fmt.Println("pod in default ns:")
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}

	deployments, err := clientset.AppsV1().Deployments("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		//handel error
		fmt.Printf("listing deployments %s\n", err.Error())
	}
	fmt.Println("deployment in default ns::")
	for _, d := range deployments.Items {
		fmt.Println(d.Name)
	}
}
