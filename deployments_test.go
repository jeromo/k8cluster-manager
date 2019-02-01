package main

import (
	"bytes"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"io"
	"io/ioutil"
	"mime/multipart"
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

func iCreateDeploymentByDescription(arg1 *gherkin.DataTable) error {
	for i := 0; i <  len(arg1.Rows); i++ {
		// arg1.Rows[i].Cells[0].Value is the nname oj the file with deploynmnt description
		bodyBuf := &bytes.Buffer{}
		bodyWriter := multipart.NewWriter(bodyBuf)

		// this step is very important
		fileWriter, err := bodyWriter.CreateFormFile("uploadfile", "test/files/" + arg1.Rows[i].Cells[0].Value)
		if err != nil {
			fmt.Println("error writing to buffer")
			return err
		}

		// open file handle
		fh, err := os.Open("test/files/" + arg1.Rows[i].Cells[0].Value)
		if err != nil {
			fmt.Println("error opening file")
			return err
		}
		defer fh.Close()

		//iocopy
		_, err = io.Copy(fileWriter, fh)
		if err != nil {
			return err
		}

		contentType := bodyWriter.FormDataContentType()
		bodyWriter.Close()

		resp, err := http.Post("http://localhost:3000/deployments/default", contentType, bodyBuf)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		resp_body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println(resp.Status)
		fmt.Println(string(resp_body))
		if resp.StatusCode != http.StatusAccepted {
			return http.ErrMissingFile
		}
		return nil

	}
	return nil
}

func iGetTheDeploymentCreated() error {
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
}


func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}





//Example Post
/*
func SendPostRequest (url string, filename string, filetype string) []byte {
    file, err := os.Open(filename)

    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()


    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    part, err := writer.CreateFormFile(filetype, filepath.Base(file.Name()))

    if err != nil {
        log.Fatal(err)
    }

    io.Copy(part, file)
    writer.Close()
    request, err := http.NewRequest("POST", url, body)

    if err != nil {
        log.Fatal(err)
    }

    request.Header.Add("Content-Type", writer.FormDataContentType())
    client := &http.Client{}

    response, err := client.Do(request)

    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    content, err := ioutil.ReadAll(response.Body)

    if err != nil {
        log.Fatal(err)
    }

    return content
}
*/
