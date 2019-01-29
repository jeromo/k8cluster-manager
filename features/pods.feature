Feature: Pods
  In order to use the cluster
  As a client
  I need to be able to manage pods

  Scenario: Get Pods
    Given the ws server is healthy running
    When I ask for pods in <namespace>
      | namespace  |
      | default    |
      | kubesystem |
      | pepito     |
    Then I get all the pods of the namespace
