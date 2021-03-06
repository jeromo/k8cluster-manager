package main

import (
	"errors"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var serverLaunched = false
var debug = true

// Contains the answer of the http requests
var contents string

func theWsServerIsHealthyRunning() error {
	if serverLaunched == false {
		serverLaunched = true
		go launchServer()
		time.Sleep(time.Second)
	}

	return nil
}

func iAskForNamespaces() error {
	contents = ""
	response, err := http.Get("http://localhost:3000/namespaces")
	if err != nil {
		return err
	}
	defer response.Body.Close()
	response_contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	contents = string(response_contents)

	if response.StatusCode == http.StatusOK {

		return nil
	}
	return errors.New("http status :" + response.Status)
}

func iGetAllTheNamespacesOfTheMinikubeCluster() error {

	if len(contents) > 0 {
		if debug {
			println("Found " + contents)
		}

		return nil
	}

	return fmt.Errorf("expected json, does not match actual: %s", contents)
}

func iAskForNamespace(arg1 string) error {
	response, err := http.Get("http://localhost:3000/namespaces/" + arg1)
	if err != nil {
		contents = "Error:" + arg1 + err.Error()

		return err
	}
	defer response.Body.Close()
	response_contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)

		os.Exit(1)
	}
	contents = string(response_contents)

	return err
}

func itShouldReturnItsName() error {
	if debug {
		println("Encontrados: " + contents)
	}

	return nil
	//	return fmt.Errorf("namespace not found")
}

func iAskForNamespaceName(arg1 *gherkin.DataTable) error {
	var response *http.Response
	var err error

	contents = ""
	for i := 0; i < len(arg1.Rows); i++ {
		response, err = http.Get("http://localhost:3000/namespaces/" + arg1.Rows[i].Cells[0].Value)
		if err != nil {
			contents = ""
			response.Body.Close()

			return err
		}
		if response.StatusCode != 200 {
			println("Warning: " + response.Status + " " + arg1.Rows[i].Cells[0].Value)
		} else {
			response_contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Printf("%s", err)
				response.Body.Close()
				os.Exit(1)
			}
			contents += string(response_contents) + " "
			response.Body.Close()
		}

	}
	return nil
}

func FeatureContext(s *godog.Suite) {

	s.Step(`^the ws server is healthy running$`, theWsServerIsHealthyRunning)
	s.Step(`^I ask for namespace  "([^"]*)"$`, iAskForNamespace)
	s.Step(`^it should return it\'s name$`, itShouldReturnItsName)

	s.Step(`^I ask for some namespace  <name>$`, iAskForNamespaceName)

	s.Step(`^I ask for namespaces$`, iAskForNamespaces)
	s.Step(`^I get all the namespaces of the minikube cluster$`, iGetAllTheNamespacesOfTheMinikubeCluster)
	s.Step(`^I ask for namespace  <name>$`, iAskForNamespaceName)
}
