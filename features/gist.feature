Feature:  Gists

    Background: Login to github
        Given I navigate to "https://github.com/login/"
        When I login with "" as email and "" as password
        Then I get the username
        
    Scenario: Create a public gist
        When I navigate to "https://gist.github.com/"
        And I write the description as "Testing create gist"
        And I write the filename as "testing_add_gist.md"
        And I write the content as "some content"
        And I click "Create public gist" button
        Then I should see gist with filename "testing_add_gist.md"

    Scenario: Update a gist
        When I navigate to "https://gist.github.com/"
        And I select one recent gist
        And I click edit gist
        And I write the description as "updated description"
        And I write the filename as "updated_testing.md"
        And I write the content as "updated content"
        And I click "Update public gist" button
        Then I should see gist with filename "updated_testing.md"

    Scenario: Delete a gist
        When I navigate to "https://gist.github.com/"
        And I select one recent gist
        And I click delete gist
        And I confirm the delete
        # Then I should not see the gist in list of gist

    Scenario: See all my gist
        When I navigate to "https://gist.github.com/"
        And I go to my gists
        Then I should see list of gist