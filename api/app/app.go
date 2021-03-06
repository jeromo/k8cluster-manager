package app

import (
	kubernetes2 "k8s.io/client-go/kubernetes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"k8cluster-manager/api/app/handler"
	"k8cluster-manager/api/app/k8manager"
)

// App has router and db instances
type App struct {
	Router    *mux.Router
	Clientset *kubernetes2.Clientset
}

// App initialize with predefined configuration
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
	a.Clientset = k8manager.Initialize()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/namespaces", a.GetNamespaces)
	a.Get("/namespaces/{name}", a.GetNamespace)
	a.Get("/pods/{namespace}", a.GetPods)
	a.Get("/deployments/{namespace}", a.GetDeployments)
	a.Post("/createdemodeployment/{namespace}", a.CreateDemoDeployment)
	a.Post("/deployments/{namespace}", a.CreateDeployment)
	a.Put("/deployments/{namespace}", a.UpdateDeployment)
	a.Delete("/deployments/{namespace}/{deployment}", a.DeleteDeployment)
	a.Delete("/deletedemodeployments/{namespace}", a.DeleteDemoDeployment)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) GetNamespaces(w http.ResponseWriter, r *http.Request) {
	handler.GetNamespaces(a.Clientset, w, r)
}

func (a *App) GetNamespace(w http.ResponseWriter, r *http.Request) {
	handler.GetNamespace(a.Clientset, w, r)
}

func (a *App) GetPods(w http.ResponseWriter, r *http.Request) {
	handler.GetPods(a.Clientset, w, r)
}

func (a *App) GetDeployments(w http.ResponseWriter, r *http.Request) {
	handler.GetDeployments(a.Clientset, w, r)
}

func (a *App) CreateDemoDeployment(w http.ResponseWriter, r *http.Request) {
	handler.CreateDemoDeployment(a.Clientset, w, r)
}

func (a *App) CreateDeployment(w http.ResponseWriter, r *http.Request) {
	handler.CreateDeploymentByYaml(a.Clientset, w, r)

}

func (a *App) UpdateDeployment(w http.ResponseWriter, r *http.Request) {
	handler.UpdateDeploymentByYaml(a.Clientset, w, r)

}

func (a *App) DeleteDemoDeployment(w http.ResponseWriter, r *http.Request) {
	handler.DeleteDemoDeployment(a.Clientset, w, r)
}

func (a *App) DeleteDeployment(w http.ResponseWriter, r *http.Request) {
	handler.DeleteDeployment(a.Clientset, w, r)
}
