package main

import (
	"bytes"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func iAskForDeploymentsInNamespace(arg1 *gherkin.DataTable) error {
	var response *http.Response
	var err error
	contents = ""

	for i := 0; i < len(arg1.Rows); i++ {
		response, err = http.Get("http://localhost:3000/deployments/" + arg1.Rows[i].Cells[0].Value)
		if err != nil {
			println("Error asking for deployments " + err.Error())

			return err
		}
		defer response.Body.Close()
		if response.StatusCode != http.StatusOK {
			contents = "Error: " + response.Status

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
	if strings.Contains(contents, "Error:") {
		println("Warning:namespace not found or namespace without deployments" + contents)
	}
	return nil
}
func iCreateDemoDeployment() error {
	contents = ""
	response, err := http.PostForm("http://localhost:3000/createdemodeployment/default", url.Values{})
	if err != nil {

		return nil
	}

	defer response.Body.Close()
	_, ioerr := ioutil.ReadAll(response.Body)
	if ioerr != nil {
		fmt.Printf("%s", ioerr)
		os.Exit(1)
	}

	if response.StatusCode != http.StatusCreated {
		contents = "Error:" + err.Error()
	}

	return err
}

func iGetItCreated() error {
	if strings.Compare(contents, "Error") == 0 {
		println("Warning! " + contents)
	}

	return nil
}

func iDeleteDemoDeployment() error {
	contents = ""
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
		contents = "Error: " + err.Error()
	}

	return nil
}

func iGetItDeletedIfExists() error {
	if strings.Compare(contents, "Error") == 0 {
		println("Warning! " + contents)
	}

	return nil
}

func iCreateDeploymentByDescription(arg1 *gherkin.DataTable) error {
	contents = ""
	for i := 0; i < len(arg1.Rows); i++ {
		// arg1.Rows[i].Cells[0].Value is the nname oj the file with deploynmnt description
		bodyBuf := &bytes.Buffer{}
		bodyWriter := multipart.NewWriter(bodyBuf)

		// this step is very important
		fileWriter, err := bodyWriter.CreateFormFile("uploadfile", "test/files/"+arg1.Rows[i].Cells[0].Value)
		if err != nil {
			contents = string("Error: writing to buffer")

			return err
		}

		// open file handle
		fh, err := os.Open("test/files/create/" + arg1.Rows[i].Cells[0].Value)
		if err != nil {
			contents = string("Error: opening file " + err.Error())

			return err
		}
		defer fh.Close()

		//iocopy
		_, err = io.Copy(fileWriter, fh)
		if err != nil {

			contents = string("Error: copying data file " + err.Error())
			return err
		}

		contentType := bodyWriter.FormDataContentType()
		bodyWriter.Close()

		resp, err := http.Post("http://localhost:3000/deployments/default", contentType, bodyBuf)
		if err != nil {
			contents = string("Error: Post call " + err.Error())

			return err
		}
		defer resp.Body.Close()
		resp_body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			contents = string("Error: ReadALl  " + err.Error())

			return err
		}
		contents = string(resp_body)

		return nil

	}
	return nil
}

func iGetTheDeploymentCreated() error {
	if strings.Contains(contents, "Error") {
		println(contents)

		return fmt.Errorf(contents)
	}

	return nil
}

func putRequest(url string, data io.Reader) (*http.Request, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, data)
	if err != nil {
		// handle error
		contents = string("Error: NewRequest " + err.Error())
		log.Fatal(err)
	}
	_, err = client.Do(req)
	if err != nil {
		// handle error
		contents = "Error: " + err.Error()
		log.Fatal(err)
	}

	return req, err

}

func iUpdateDeploymentByDescription(arg1 *gherkin.DataTable) error {
	contents = ""
	for i := 0; i < len(arg1.Rows); i++ {
		// arg1.Rows[i].Cells[0].Value is the nname oj the file with deploynmnt description
		bodyBuf := &bytes.Buffer{}
		bodyWriter := multipart.NewWriter(bodyBuf)

		// this step is very important
		fileWriter, err := bodyWriter.CreateFormFile("uploadfile", "test/files/update"+arg1.Rows[i].Cells[0].Value)
		if err != nil {
			fmt.Println("error writing to buffer")
			contents = string("Error: writing to buffer " + err.Error())

			return err
		}

		// open file handle
		fh, err := os.Open("test/files/update/" + arg1.Rows[i].Cells[0].Value)
		if err != nil {
			fmt.Println("error opening file " + err.Error())
			contents = string("Error: opening file " + err.Error())

			return err
		}
		defer fh.Close()

		//iocopy
		_, err = io.Copy(fileWriter, fh)
		if err != nil {

			fmt.Println("error copying data file")
			contents = string("Error: copying data file " + err.Error())

			return err
		}

		//		contentType := bodyWriter.FormDataContentType()
		bodyWriter.Close()

		req, err := putRequest("http://localhost:3000/deployments/default", bodyBuf)
		if err != nil {
			contents = string("Error: putRequest " + err.Error())

			return err
		}

		defer req.Body.Close()
		req_body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println("error reading Body")
			contents = string("Error: reading Body" + err.Error())

			return err
		}
		contents = string(req_body)

		return nil

	}
	return nil
}

func iDeleteDeploymentByName(arg1 *gherkin.DataTable) error {
	return godog.ErrPending
}

func iGetTheDeploymentDeleted() error {
	return godog.ErrPending

}
func FeatureDeploymentsContext(s *godog.Suite) {
	s.Step(`^I ask for deployments in <namespace>$`, iAskForDeploymentsInNamespace)
	s.Step(`^I get all the deployments of the namespace$`, iGetAllTheDeploymentsOfTheNamespace)

	s.Step(`^I create  demo deployment$`, iCreateDemoDeployment)
	s.Step(`^I get it created$`, iGetItCreated)

	s.Step(`^I delete  demo deployment$`, iDeleteDemoDeployment)
	s.Step(`^I get it deleted if exists$`, iGetItDeletedIfExists)

	s.Step(`^I create deployment by <description>$`, iCreateDeploymentByDescription)
	s.Step(`^I get the deployment created$`, iGetTheDeploymentCreated)

	s.Step(`^I update deployment by <description>$`, iUpdateDeploymentByDescription)

	s.Step(`^I delete deployment by <name>$`, iDeleteDeploymentByName)
	s.Step(`^I get the deployment deleted$`, iGetTheDeploymentDeleted)
}
