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
			contents += "Warning: " + response.Status + " " + arg1.Rows[i].Cells[0].Value + " "
		    //A HTTP status different than StatusOK is not considered an error by now
		}
		response_contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)

			os.Exit(1)
		}
		if response.StatusCode == http.StatusOK {
			if response_contents != nil {
				contents += string(response_contents)
			}
		}

	}

	return nil
}

func iGetAllTheDeploymentsOfTheNamespace() error {
	if debug {
		println("Deployments " + contents)
	}

	return nil
}

func iCreateDemoDeployment() error {
	contents = ""
	response, err := http.PostForm("http://localhost:3000/createdemodeployment/default", url.Values{})
	if err != nil {

		return err
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
			contents = string("Error: CreateFormFile")

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
		contents += string(resp_body) + " "
	}
	return nil
}

func iGetTheDeploymentCreated() error {
	if strings.Contains(contents, "Error") {
		return fmt.Errorf(contents)
	}

	if debug {
		println(contents)
	}
	return nil
}

func putRequest(url string, contentType string, data io.Reader) (*http.Response, *http.Request, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, data)
	if err != nil {
		contents = string("Error: NewRequest " + err.Error())

		log.Fatal(err)
	}
	req.Header.Set("Content-Type", contentType)
	resp, err := client.Do(req)
	if err != nil {

		log.Fatal(err)
	}

	return resp, req, err

}

func iUpdateDeploymentByDescription(arg1 *gherkin.DataTable) error {
	contents = ""
	for i := 0; i < len(arg1.Rows); i++ {
		// arg1.Rows[i].Cells[0].Value is the nname oj the file with deploynmnt description
		bodyBuf := &bytes.Buffer{}
		bodyWriter := multipart.NewWriter(bodyBuf)

		fileWriter, err := bodyWriter.CreateFormFile("uploadfile", "test/files/update"+arg1.Rows[i].Cells[0].Value)
		if err != nil {
			contents = string("Error: writing to buffer " + err.Error())

			return err
		}

		// open file handle
		fh, err := os.Open("test/files/update/" + arg1.Rows[i].Cells[0].Value)
		if err != nil {
			contents = string("Error: opening file " + err.Error())

			return err
		}
		defer fh.Close()

		_, err = io.Copy(fileWriter, fh)
		if err != nil {

			contents = string("Error: copying data file " + err.Error())

			return err
		}

		contentType := bodyWriter.FormDataContentType()
		bodyWriter.Close()

		resp, req, err := putRequest("http://localhost:3000/deployments/default", contentType, bodyBuf)

		if err != nil {
			contents = string("Error: putRequest " + err.Error())

			return err
		}

		defer req.Body.Close()
		req_body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			contents = string("Error: reading Body" + err.Error())

			return err
		}
		contents += string(req_body) + " "


		if http.StatusOK != resp.StatusCode {
			contents = "Error: "+ resp.Status

			return fmt.Errorf("Error %s", resp.Status)
		}

	}

	return nil
}

func iDeleteDeploymentByName(arg1 *gherkin.DataTable) error {
	contents = ""
	for i := 0; i < len(arg1.Rows); i++ {
		// arg1.Rows[i].Cells[0].Value is the nname oj the file with deploynmnt description
		req, err := http.NewRequest(
			"DELETE", "http://localhost:3000/deployments/default/"+arg1.Rows[i].Cells[0].Value, nil)
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
			contents += "Error: " + arg1.Rows[i].Cells[0].Value + " " + err.Error() + " "
		} else {
			contents += arg1.Rows[i].Cells[0].Value + "deleted "
		}
	}

	return nil
}

func iGetTheDeploymentDeleted() error {
	if strings.Compare(contents, "Error") == 0 {

		return  fmt.Errorf(contents)
	}

	if debug {
		println(contents)
	}

	return nil
}

func iGetTheDeploymentUpdated() error {
	if strings.Compare(contents, "Error") == 0 {

		return  fmt.Errorf(contents)
	}

	if debug {
		println(contents)
	}

	return nil
}
func iCreateDeploymentByBadDescription(arg1 *gherkin.DataTable) error {
	err := iCreateDeploymentByDescription(arg1)
	if err == nil {

		return fmt.Errorf("Should return an error")
	}

	if debug {
		println(contents)
	}

	return nil
}

func iReceiveCreateError() error {
	iGetTheDeploymentCreated()
	if !strings.Contains(contents, "Error") {

		return fmt.Errorf("Should return an error")
	}

	return nil
}

func iUpdateDeploymentByNonExistentDescription(arg1 *gherkin.DataTable) error {
	err := iUpdateDeploymentByDescription(arg1)
	if err == nil {

		return fmt.Errorf("Should return an error")
	}

	if debug {
		println(contents)
	}

	return nil
}

func iGetUpdateError() error {
	if !strings.Contains(contents, "Error") {

		return fmt.Errorf("Should return an error")
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

	s.Step(`^I create deployment by <description>$`, iCreateDeploymentByDescription)
	s.Step(`^I get the deployment created$`, iGetTheDeploymentCreated)

	s.Step(`^I update deployment by <description>$`, iUpdateDeploymentByDescription)
	s.Step(`^I get the deployment updated$`, iGetTheDeploymentUpdated)

	s.Step(`^I delete deployment by <name>$`, iDeleteDeploymentByName)
	s.Step(`^I get the deployment deleted$`, iGetTheDeploymentDeleted)

	s.Step(`^I create deployment by bad <description>$`, iCreateDeploymentByBadDescription)
	s.Step(`^I receive create error$`, iReceiveCreateError)

	s.Step(`^I update deployment by non existent <description>$`, iUpdateDeploymentByNonExistentDescription)
	s.Step(`^I get update error$`, iGetUpdateError)
}
