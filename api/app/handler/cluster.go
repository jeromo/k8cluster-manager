package handler

import (
	"k8cluster-manager/api/app/k8manager"
	kubernetes2 "k8s.io/client-go/kubernetes"
	"net/http"
)
 
func GetNamespaces(clientset *kubernetes2.Clientset, w http.ResponseWriter, r *http.Request) {
	salida := k8manager.GetNamespaces(clientset)
	respondJSON(w, http.StatusOK, salida)
}

func GetNamespace(name string, clientset *kubernetes2.Clientset, w http.ResponseWriter, r *http.Request) {
	salida := k8manager.GetNamespace(name, clientset)
	respondJSON(w, http.StatusOK, salida)
}
