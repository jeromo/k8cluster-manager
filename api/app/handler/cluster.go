package handler

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"k8cluster-manager/api/app/k8manager"
	"k8s.io/apimachinery/pkg/api/errors"
	kubernetes2 "k8s.io/client-go/kubernetes"
	"net/http"
)

func GetNamespaces(clientset *kubernetes2.Clientset, w http.ResponseWriter, r *http.Request) {
	output, err := k8manager.GetNamespaces(clientset)
	if err == nil {
		respondJSON(w, http.StatusOK, output)
	} else {
		respondJSON(w, http.StatusInternalServerError, output)
	}
}

func GetNamespace(name string, clientset *kubernetes2.Clientset, w http.ResponseWriter, r *http.Request) {
	output, err := k8manager.GetNamespace(name, clientset)
	if err == nil {
		respondJSON(w, http.StatusOK, output)
	} else {
		respondJSON(w, http.StatusNotFound, output)
	}
}

func GetPods(namespace string, clientset *kubernetes2.Clientset, w http.ResponseWriter, r *http.Request) {
	output, err := k8manager.GetPods(namespace, clientset)
	if err == nil {
		respondJSON(w, http.StatusOK, output)
	} else {
		respondJSON(w, http.StatusNotFound, output)
	}
}

func GetDeployments(namespace string, clientset *kubernetes2.Clientset, w http.ResponseWriter, r *http.Request) {
	output, err := k8manager.GetDeployments(namespace, clientset)
	if err != nil {
		respondJSON(w, http.StatusNotFound, output)
	} else {
		respondJSON(w, http.StatusOK, output)
	}
}

func CreateDemoDeployment(clientset *kubernetes2.Clientset, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	namespace := vars["namespace"]

	output, err := k8manager.CreateDemoDeployment(namespace, clientset)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, output)
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

	output, err := k8manager.CreateDeploymentByYaml(namespace, body, clientset)
	if err != nil {
		println("CreateDeploymentByYaml returns :" + err.Error())
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

	output, err := k8manager.UpdateDeploymentByYaml(namespace, body, clientset)
	if err != nil {
		println("UpdateDeploymentByYaml returns :" + err.Error())
		respondJSON(w, http.StatusConflict, output)
	} else {
		respondJSON(w, http.StatusAccepted, output)
	}
}

func DeleteDemoDeployment(namespace string, clientset *kubernetes2.Clientset, w http.ResponseWriter, r *http.Request) {
	output, err := k8manager.DeleteDemoDeployment(namespace, clientset)
	if err == nil {
		respondJSON(w, http.StatusAccepted, output)
	} else {
		println(err.Error())
		if errors.IsNotFound(err) {
			respondJSON(w, http.StatusNotFound, output)
		} else {
			respondJSON(w, http.StatusInternalServerError, namespace)
		}
	}
}

func DeleteDeployment(namespace string, clientset *kubernetes2.Clientset, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deployment := vars["deployment"]

	output, err := k8manager.DeleteDeployment(deployment, clientset)
	if err != nil {
		respondJSON(w, http.StatusNotFound, namespace)
	} else {
		respondJSON(w, http.StatusAccepted, output)
	}
}
