package k8manager

import (
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes2 "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/util/retry"
)

func GetDeployments(namespace string, clientset *kubernetes2.Clientset) ([]string, error) {
	deploymentsClient := clientset.AppsV1().Deployments(namespace)
	var output []string
	list, err := deploymentsClient.List(metav1.ListOptions{})
	if err != nil {
		return output, err
	}
	for _, d := range list.Items {
		output = append(output, d.Name+" ")
	}

	return output, err
}

func CreateDemoDeployment(namespace string, clientset *kubernetes2.Clientset) (string, error) {
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

		return "", err

	}

	return result.GetObjectMeta().GetName(), err
}

func int32Ptr(i int32) *int32 { return &i }

func DeleteDemoDeployment(namespace string, clientset *kubernetes2.Clientset) (string, error) {
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	deletePolicy := metav1.DeletePropagationForeground
	err := deploymentsClient.Delete("demo-deployment", &metav1.DeleteOptions{PropagationPolicy: &deletePolicy})
	if err != nil && !errors.IsNotFound(err) {
		panic(err)
	}
	if errors.IsNotFound(err) {

		return "NotFound", err
	}

	return "Deleted", err
}

func CreateDeploymentByYaml(namespace string, configFile []byte, clientset *kubernetes2.Clientset) (string, error) {
	//var deployment appsv1.Deployment

	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, err := decode([]byte(configFile), nil, nil)

	if err != nil {

		return "", err
	}

	// Create Deployment
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	switch o := obj.(type) {
	case *appsv1.Deployment:
		result, err := deploymentsClient.Create(o)
		if err != nil {

			return "", err
		}

		return result.GetObjectMeta().GetName(), err
	default:
		//o is unknown for us
		err = errors.NewBadRequest("Object type described in yaml not found")

		return "", err
	}
}

func UpdateDeploymentByYaml(namespace string, configFile []byte, clientset *kubernetes2.Clientset) (string, error) {
	//var deployment appsv1.Deployment

	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, err := decode([]byte(configFile), nil, nil)

	if err != nil {

		return "", err
	}

	// Update Deployment
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	switch o := obj.(type) {
	case *appsv1.Deployment:
		retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
			// Retrieve the latest version of Deployment before attempting update
			// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
			_, updateErr := deploymentsClient.Update(o)

			return updateErr
		})
		if retryErr != nil {
			return "", retryErr
		}
		return o.GetObjectMeta().GetName(), err
	default:
		//o is unknown for us
		err = errors.NewBadRequest("Object type described in yaml not found")

		return "", err
	}
}

func DeleteDeployment(deployment string, clientset *kubernetes2.Clientset) (string, error) {
	deletePolicy := metav1.DeletePropagationForeground

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	if err := deploymentsClient.Delete(deployment, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {

		return "", err
	}

	return deployment, nil
}
