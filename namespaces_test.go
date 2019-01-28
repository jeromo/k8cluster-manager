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
	if (!serverLaunched){
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

func FeatureNamespacesContext(s *godog.Suite) {
	s.Step(`^I ask for namespaces$`, iAskForNamespaces)
	s.Step(`^I get all the namespaces of the minikube cluster$`, iGetAllTheNamespacesOfTheMinikubeCluster)
}

func iAskForNamespaceString(arg1 *gherkin.DataTable) error {
	return godog.ErrPending
}

func thereShouldReturnItsName() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	api := &apiFeature{}

	s.BeforeScenario(api.resetResponse)

	s.Step(`^the ws server is healthy running$`, theWsServerIsHealthyRunning)
	s.Step(`^I ask for namespace <string>$`, iAskForNamespaceString)
	s.Step(`^there should return it\'s name$`, thereShouldReturnItsName)
}

