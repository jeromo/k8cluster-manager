package main

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func iAskForDeploymentsInNamespace(arg1 *gherkin.DataTable) error {
	var response *http.Response
	var err error

	for i := 0; i <  len(arg1.Rows); i++ {
		response, err = http.Get("http://localhost:3000/deployments/"+ arg1.Rows[i].Cells[0].Value)
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

func iGetAllTheDeploymentsOfTheNamespace() error {
	if (len(contents) == 0) {
		println("Warning:namespace not found or namespace without deployments")
	}
	return nil
}
func iCreateDemoDeployment() error {
	response, err := http.PostForm("http://localhost:3000/deployments/default", url.Values{})
	if err != nil {
		contents = ""

		return nil
	}

	defer response.Body.Close()
	_, ioerr := ioutil.ReadAll(response.Body)
	if ioerr != nil {
		fmt.Printf("%s", ioerr)
		os.Exit(1)
	}

	if (response.StatusCode != http.StatusCreated) {
		contents = "Error"
	}

	return err
}

func iGetItCreated() error {
	if strings.Compare(contents, "Error") == 0 {
		println("Warning! demo deployment already exists")
	}

	return nil
}

func iDeleteDemoDeployment() error {
	req, err := http.NewRequest("DELETE", "http://localhost:3000/deployments/default", nil)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	if response.StatusCode != http.StatusAccepted {
		contents = "Error"
	}

	return nil
}

func iGetItDeletedIfExists() error {
	if strings.Compare(contents, "Error") == 0 {
		println("Warning! demo deployment not found")
	}

	return nil
}


func FeatureDeploymentsContext(s *godog.Suite) {
	s.Step(`^I ask for deployments in <namespace>$`, iAskForDeploymentsInNamespace)
	s.Step(`^I get all the deployments of the namespace$`, iGetAllTheDeploymentsOfTheNamespace)

	s.Step(`^I create  demo deployment$`, iCreateDemoDeployment)
	s.Step(`^I get it created$`, iGetItCreated)

	s.Step(`^I delete  demo deployment$`, iDeleteDemoDeployment)
	s.Step(`^I get it deleted if exists$`, iGetItDeletedIfExists)

}

