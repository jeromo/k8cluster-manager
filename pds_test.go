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

	for i := 0; i <  len(arg1.Rows); i++ {
		response, err = http.Get("http://localhost:3000/pods/"+ arg1.Rows[i].Cells[0].Value)
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

func iGetAllThePodsOfTheNamespace() error {
	if (len(contents) == 0) {
		println("Warning:namespace not found or namespace without pods")
	}
	return nil
}

func FeaturePodsContext(s *godog.Suite) {
	s.Step(`^I ask for pods in <namespace>$`, iAskForPodsInNamespace)
	s.Step(`^I get all the pods of the namespace$`, iGetAllThePodsOfTheNamespace)
}
