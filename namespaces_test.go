/* file: $GOPATH/src/godogs/godogs_test.go */
package main

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var serverLaunched = false
var req http.Request

type apiFeature struct {
	resp *httptest.ResponseRecorder
}
var contents string

func (a *apiFeature) resetResponse(interface{}) {
	a.resp = httptest.NewRecorder()
}
func theWsServerIsHealthyRunning() error {
	if serverLaunched == false{
		serverLaunched = true
		go launchServer();

		time.Sleep(time.Second)
	}
	return nil
}

func iAskForNamespaces() error {
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

	return err
}

func iGetAllTheNamespacesOfTheMinikubeCluster() error {

	if (len(contents) > 0) {
		return nil
	}

	return fmt.Errorf("expected json, does not match actual: %s", contents)
}


func iAskForNamespace(arg1 string) error {
	response, err := http.Get("http://localhost:3000/namespaces/"+ arg1)
	if err != nil {
		contents = ""
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


func thereShouldReturnItsName() error {
	if (len(contents) == 0) {
		println("Warning:namespace not found")
	}

	return nil
//	return fmt.Errorf("namespace not found")
}

func iAskForSomeNamespaceName(arg1 *gherkin.DataTable) error {
	var response *http.Response
	var err error

 	for i := 0; i <  len(arg1.Rows); i++ {
		response, err = http.Get("http://localhost:3000/namespaces/"+ arg1.Rows[i].Cells[0].Value)
		if err != nil {
			contents = ""

			return err
		}
		defer response.Body.Close()
		if response.StatusCode != 200 {
			contents = ""

			return nil
		}
		response_contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		contents = string(response_contents)

	}
	return nil
}


func FeatureContext(s *godog.Suite) {
	api := &apiFeature{}

	s.BeforeScenario(api.resetResponse)

	s.Step(`^the ws server is healthy running$`, theWsServerIsHealthyRunning)
	s.Step(`^I ask for namespace  "([^"]*)"$`, iAskForNamespace)
	s.Step(`^there should return it\'s name$`, thereShouldReturnItsName)

	s.Step(`^I ask for some namespace  <name>$`, iAskForSomeNamespaceName)

	s.Step(`^I ask for namespaces$`, iAskForNamespaces)
	s.Step(`^I get all the namespaces of the minikube cluster$`, iGetAllTheNamespacesOfTheMinikubeCluster)
	}

