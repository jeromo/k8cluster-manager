package k8manager

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes2 "k8s.io/client-go/kubernetes"
)

func GetPods(namespace string, clientset *kubernetes2.Clientset) ([]string, error) {
	var output []string
	pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {

		return output, err
	}
	for i := 0; i < len(pods.Items); i++ {
		output = append(output, pods.Items[i].ObjectMeta.Name+" ")
	}

	return output, err
}
