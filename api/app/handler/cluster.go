package handler

import (
	"k8cluster-manager/api/app/k8manager"
	kubernetes2 "k8s.io/client-go/kubernetes"
	"net/http"
	"strings"
)
 
func GetNamespaces(clientset *kubernetes2.Clientset, w http.ResponseWriter, r *http.Request) {
	salida := k8manager.GetNamespaces(clientset)
	respondJSON(w, http.StatusOK, salida)
}

func GetNamespace(name string, clientset *kubernetes2.Clientset, w http.ResponseWriter, r *http.Request) {
	output := k8manager.GetNamespace(name, clientset)
	if strings.Compare(output,"Not found") == 0 {
		respondJSON(w, http.StatusNotFound, name)
	} else {
		respondJSON(w, http.StatusOK, output)
	}
}


func GetPods(namespace string, clientset *kubernetes2.Clientset, w http.ResponseWriter, r *http.Request) {
	output := k8manager.GetPods(namespace, clientset)
	if output == nil {
		respondJSON(w, http.StatusNotFound, namespace)
	} else {
		respondJSON(w, http.StatusOK, output)
	}
}


func GetDeployments(namespace string, clientset *kubernetes2.Clientset, w http.ResponseWriter, r *http.Request) {
	output := k8manager.GetDeployments(namespace, clientset)
	if output == nil {
		respondJSON(w, http.StatusNotFound, namespace)
	} else {
		respondJSON(w, http.StatusOK, output)
	}
}


func CreateDemoDeployment(namespace string, clientset *kubernetes2.Clientset, w http.ResponseWriter, r *http.Request) {
	output := k8manager.CreateDemoDeployment(namespace, clientset)
	if strings.Compare(output,"Error") == 0 {
		respondJSON(w, http.StatusInternalServerError, namespace)
	} else {
		respondJSON(w, http.StatusAccepted, output)
	}
}


func DeleteDemoDeployment(namespace string, clientset *kubernetes2.Clientset, w http.ResponseWriter, r *http.Request) {
	output := k8manager.DeleteDemoDeployment(namespace, clientset)
	if strings.Compare(output,"NotFound") == 0 {
		respondJSON(w, http.StatusNotFound, namespace)
	} else {
		respondJSON(w, http.StatusAccepted, output)
	}
}

