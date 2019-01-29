package k8manager

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes2 "k8s.io/client-go/kubernetes"
)

func GetDeployments( namespace string, clientset *kubernetes2.Clientset) []string {
	deploymentsClient := clientset.AppsV1().Deployments(namespace)
	fmt.Printf("Listing deployments in namespace %q:\n", namespace)
	list, err := deploymentsClient.List(metav1.ListOptions{})
	if err != nil {
		return nil
	}
	var output []string
	for _, d := range list.Items {
		output = append(output,d.Name)
	}

	return output
}
