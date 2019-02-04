package k8manager

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes2 "k8s.io/client-go/kubernetes"
)

func GetNamespaces( clientset *kubernetes2.Clientset) []string{
	namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d namespacess in the cluster\n", len(namespaces.Items))

	var output []string
	for i :=0; i < len(namespaces.Items); i++ {
		output = append(output, namespaces.Items[i].ObjectMeta.Name +" ")
	}


	return output
}

func GetNamespace( name string, clientset *kubernetes2.Clientset) string{
	namespace, err := clientset.CoreV1().Namespaces().Get(name, metav1.GetOptions{})
	if err != nil {
		return "Error: " + err.Error()
	}

	return 	namespace.ObjectMeta.Name
}

