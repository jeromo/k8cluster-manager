package main

import (
	"k8cluster-manager/api/app"
	"os"
	"sync"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

// SafeCounter is safe to use concurrently.
type SafeState struct {
	stop bool
	mux  sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeState) Stop() {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.stop = true
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeState) WantStop() bool {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.stop
}

var State SafeState

func main() {
	launchServer()
}

func launchServer() {
	app := &app.App{}
	app.Initialize()
	app.Run(":3000")
}

/*func launchServer() {
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

	//Seek devFAAS namespacxe, if doesn't exist, create it
	namespace, err := clientset.CoreV1().Namespaces().Get("devFAAS", metav1.GetOptions{})
	println("Err is")
	println(err)
	if errors.IsNotFound(err) {
		fmt.Printf("Namespace devFaas not found\n")
	} else {
		fmt.Printf("Namespace %s found\n", namespace.Name)

	}
	for {
		if State.WantStop() {
			return
		}
		namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d namespacess in the cluster\n", len(namespaces.Items))

		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions like e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		namespace := "default"
		//		pod := "kubernetes-bootcamp-7799cbcb86-j2q8q"
		pod := "kubernetes-bootcamp"
		_, err = clientset.CoreV1().Pods(namespace).Get(pod, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %s in namespace %s: %v\n",
				pod, namespace, statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
		}

		time.Sleep(10 * time.Second)
	}
}
*/

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
