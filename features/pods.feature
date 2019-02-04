Feature: Pods
  In order to use the cluster
  As a client
  I need to be able to manage pods

  Scenario: Get Pods
    Given the ws server is healthy running
    When I ask for pods in <namespace>
      | kube-public  |
      | default    |
      | kube-system |
      | pepito     |
    Then I get all the pods
