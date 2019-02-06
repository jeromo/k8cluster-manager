Feature: Deployments
  In order to use the cluster
  As a client
  I need to be able to manage deployments

  Scenario: Get Deployments
    Given the ws server is healthy running
    When I ask for deployments in <namespace>
      | default    |
      | kube-system |
      | pepito     |
    Then I get all the deployments of the namespace

  Scenario: Create deployment by file
    Given the ws server is healthy running
    When I create deployment by <description>
      | redis-master.yaml |
    Then I get it created

  Scenario: Create deployment by bad file
    Given the ws server is healthy running
    When I create deployment by bad <description>
      | error.yaml |
    Then I receive create error

  Scenario: Update deployment by file
    Given the ws server is healthy running
    When I update deployment by <description>
      | redis-master.yaml |
    Then I get the deployment updated

  Scenario: Update non existent deployment by file
    Given the ws server is healthy running
    When I update deployment by non existent <description>
      | redis-slave-deployment.yaml |
    Then I get update error

  Scenario: Delete deployment by name
    Given the ws server is healthy running
    When I delete deployment by <name>
      | redis-master |
    Then I get the deployment deleted
