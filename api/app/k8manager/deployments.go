package k8manager

import (
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes2 "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
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


func CreateDemoDeployment( namespace string, clientset *kubernetes2.Clientset) string {
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	// Create Deployment
	result, err := deploymentsClient.Create(deployment)
	if err != nil {
		return "Error"

	}
	return result.GetObjectMeta().GetName()
}

func int32Ptr(i int32) *int32 { return &i }

func DeleteDemoDeployment( namespace string, clientset *kubernetes2.Clientset) string {
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	deletePolicy := metav1.DeletePropagationForeground
	err := deploymentsClient.Delete("demo-deployment", &metav1.DeleteOptions{PropagationPolicy: &deletePolicy,})
	if err != nil && !errors.IsNotFound(err){
		panic(err)
	}
	if errors.IsNotFound(err) {

		return "NotFound"
	}

	return "Deleted"
}

func CreateDeploymentByYaml( namespace string, configFile []byte, clientset *kubernetes2.Clientset) string {
	//var deployment appsv1.Deployment

	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, err := decode([]byte(configFile), nil, nil)

	if err != nil {
		return "Error"
	}

	// Create Deployment
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	switch o := obj.(type) {
	case *appsv1.Deployment:
		result, err := deploymentsClient.Create(o)
		return result.GetObjectMeta().GetName()
		if err != nil {

			return "Error"
		}
	default:
		//o is unknown for us
	}
	return "Error"
}
