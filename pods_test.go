/* file: $GOPATH/src/godogs/godogs_test.go */
package main

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"io/ioutil"
	"net/http"
	"os"
)

func iAskForPodsInNamespace(arg1 *gherkin.DataTable) error {
	var response *http.Response
	var err error

	contents = ""
	for i := 0; i < len(arg1.Rows); i++ {
		response, err = http.Get("http://localhost:3000/pods/" + arg1.Rows[i].Cells[0].Value)
		if err != nil {
			contents = ""
			response.Body.Close()

			return err
		}
		if response.StatusCode != http.StatusOK {
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

func iGetAllThePods() error {
	println("Pods: " + contents)
	return nil
}

func FeaturePodsContext(s *godog.Suite) {
	s.Step(`^I ask for pods in <namespace>$`, iAskForPodsInNamespace)
	s.Step(`^I get all the pods$`, iGetAllThePods)
}
