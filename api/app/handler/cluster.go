package handler

import (
	"k8cluster-manager/api/app/k8manager"
	kubernetes2 "k8s.io/client-go/kubernetes"
	"net/http"
)
 
func GetNamespaces(clientset *kubernetes2.Clientset, w http.ResponseWriter, r *http.Request) {
	salida := k8manager.GetNamespaces(clientset)
	respondJSON(w, http.StatusOK, salida)
/*	employees := []model.Employee{}
	db.Find(&employees)
	respondJSON(w, http.StatusOK, employees)
*/}
 
