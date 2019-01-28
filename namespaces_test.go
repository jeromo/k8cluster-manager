/* file: $GOPATH/src/godogs/godogs_test.go */
package main

import (
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
)

func theWsServerIsHealthyRunning() error {
	go launchServer();
	return nil
}

func iAskForNamespaceString(arg1 *gherkin.DataTable) error {
	return godog.ErrPending
}

func thereShouldReturnItsName() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^the ws server is healthy running$`, theWsServerIsHealthyRunning)
	s.Step(`^I ask for namespace <string>$`, iAskForNamespaceString)
	s.Step(`^there should return it\'s name$`, thereShouldReturnItsName)
}

