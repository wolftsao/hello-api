Feature: Translation Service
  Users should be able to submit a word to translate words within the application

  @smoke-test
  Scenario: Translation
    Given the word "hello"
    When I translate it to "german"
    Then the response should be "hallo"

  @smoke-test
  Scenario: Translation unknown
    Given the word "goodbye"
    When I translate it to "german"
    Then the response should be ""

  @smoke-test
  Scenario: Translation
    Given the word "hello"
    When I translate it to "bulgarian"
    Then the response should be "Здравейте"

  @regression-test
  Scenario: Translation Czech
    Given the word "hello"
    When I translate it to "Czech"
    Then the response should be "Ahoj"