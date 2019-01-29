package k8manager

import (
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	kubernetes2 "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func Initialize() *kubernetes2.Clientset{
	var kubeconfig *string
	if home := homeDir(); home != "" {
	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
	panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
	panic(err.Error())
	}

	return clientset
}

func GetNamespaces( clientset *kubernetes2.Clientset) []string{
	namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d namespacess in the cluster\n", len(namespaces.Items))

	var output []string
	for i :=0; i < len(namespaces.Items); i++ {
		output = append(output, namespaces.Items[i].ObjectMeta.Name)
	}


	return output
}

func GetNamespace( name string, clientset *kubernetes2.Clientset) string{
	namespace, err := clientset.CoreV1().Namespaces().Get(name, metav1.GetOptions{})
	if err != nil {
		return "Not found"
	}

	return 	namespace.ObjectMeta.Name
}

func GetPods( namespace string, clientset *kubernetes2.Clientset) []string {
	pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil
	}
	var output []string
	for i :=0; i < len(pods.Items); i++ {
		output = append(output, pods.Items[i].ObjectMeta.Name)
	}

	return output
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

