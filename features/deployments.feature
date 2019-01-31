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

  Scenario: Create demo deployment
    Given the ws server is healthy running
    When I create  demo deployment
    Then I get it created

  Scenario: Delete demo deployment
    Given the ws server is healthy running
    When I delete  demo deployment
    Then I get it deleted if exists

  Scenario: Create deployment by file
    Given the ws server is healthy running
    When I create deployment by <description>
      | description           |
      | nginx-deployment.yaml |
    Then I get the deployments created

