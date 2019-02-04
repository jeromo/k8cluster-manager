package handler

import (
	"github.com/gorilla/mux"
	"io/ioutil"
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


func CreateDemoDeployment(clientset *kubernetes2.Clientset, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	namespace := vars["namespace"]

	output := k8manager.CreateDemoDeployment(namespace, clientset)
	if strings.Compare(output,"Error") == 0 {
		respondJSON(w, http.StatusInternalServerError, namespace)
	} else {
		respondJSON(w, http.StatusAccepted, output)
	}
}

func CreateDeploymentByYaml(clientset *kubernetes2.Clientset, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	namespace := vars["namespace"]

	file, _, err := r.FormFile("uploadfile")
	if err != nil {
		respondJSON(w, http.StatusNotAcceptable, namespace)

		return
	}
	defer file.Close()

	body, err := ioutil.ReadAll(file)

	if err != nil {
		respondJSON(w, http.StatusBadRequest, namespace)

		return
	}

	output := k8manager.CreateDeploymentByYaml(namespace, body, clientset)
	if strings.HasPrefix(output,"Error") {
		respondJSON(w, http.StatusConflict, output)
	} else {
		respondJSON(w, http.StatusAccepted, output)
	}
}

func UpdateDeploymentByYaml(clientset *kubernetes2.Clientset, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	namespace := vars["namespace"]

	file, _, err := r.FormFile("uploadfile")
	if err != nil {
		respondJSON(w, http.StatusNotAcceptable, namespace)

		return
	}
	defer file.Close()

	body, err := ioutil.ReadAll(file)

	if err != nil {
		respondJSON(w, http.StatusBadRequest, namespace)

		return
	}

	output := k8manager.UpdateDeploymentByYaml(namespace, body, clientset)
	if strings.HasPrefix(output,"Error") {
		respondJSON(w, http.StatusConflict, output)
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

