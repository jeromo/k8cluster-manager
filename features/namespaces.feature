Feature: Namespaces
  In order to use the cluster
  As a client
  I need to be able to manage namespaces

  Scenario: Get Namespaces
    Given the ws server is healthy running
    When I ask for namespaces
    Then I get all the namespaces of the minikube cluster

  Scenario: Get Namespace
    Given the ws server is healthy running
    When I ask for namespace  "default"
    Then it should return it's name


  Scenario: Get Some Namespaces
    Given the ws server is healthy running
    When I ask for namespace  <name>
    | default  |
    | devFAAS  |
    Then it should return it's name
