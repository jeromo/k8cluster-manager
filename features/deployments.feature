Feature: Deployments
  In order to use the cluster
  As a client
  I need to be able to manage deployments

  Scenario: Get Deployments
    Given the ws server is healthy running
    When I ask for deployments in <namespace>
      | namespace  |
      | default    |
      | kubesystem |
      | pepito     |
    Then I get all the deployments of the namespace
