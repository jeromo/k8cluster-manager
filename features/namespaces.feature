Feature: Namspaces
  In order to be happy
  As a hungry gopher
  I need to be able to eat godogs

  Scenario: Get Namespaces
    Given the ws server is healthy running
    When I ask for namespaces
    Then I get all the namespaces of the minikube cluster

  Scenario: Get Namespace devFAAS
Given the ws server is healthy running
When I ask for namespace <string>
| string  |
| DevFAAS |
Then there should return it's name
