Feature: Search package on pkg.go.dev

  Scenario: Navigate to packages path
    Given I visit "https://go.dev"
    And I can see the logo
    When I click packages on the menu
    Then I should be redirected to "https://pkg.go.dev"

  Scenario Outline: Search package on https://pkg.go.dev
    Given I visit "https://go.dev"
    And I navigate to "https://pkg.go.dev" by clicking packages on menu
    When I enter "<package>" package name in the search
    When I press search button
    Then I should see a search page with "Godog" package

    Examples:
      | package |
      | godog    |
      | dotenv   |