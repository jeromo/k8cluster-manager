Feature: Namspaces
  In order to be happy
  As a hungry gopher
  I need to be able to eat godogs

Scenario: Get Namespace devFAAS
Given the ws server is healthy running
When I ask for namespace <string>
| string  |
| DevFAAS |
Then there should return it's name
