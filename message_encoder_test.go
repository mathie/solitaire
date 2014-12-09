package main

import "testing"

func TestAlreadyFormattedString(t *testing.T) {
  message := "HELLO WORLD"
  expected := []string{ "HELLO", "WORLD" }
  assertEncoded(expected, message, t)
}

func TestCovertsToUpperCase(t *testing.T) {
 message := "Hello World"
 expected := []string{ "HELLO", "WORLD" }
 assertEncoded(expected, message, t)
}

func TestRemovesNonAlphaCharacters(t *testing.T) {
  message := "Hello, World!"
  expected := []string{ "HELLO", "WORLD" }
  assertEncoded(expected, message, t)
}

func TestPadding(t *testing.T) {
  message := "Goodbye, cruel world!"
  expected := []string{ "GOODB", "YECRU", "ELWOR", "LDXXX" }
  assertEncoded(expected, message, t)
}

func TestDifferentBlockSize(t *testing.T) {
  message := "Goodbye, cruel world!"
  expected := []string{ "GOODBYECRU", "ELWORLDXXX" }

  actual := EncodeMessageWithBlockSize(message, 10)

  assertSameStringArray(expected, actual, t)
}

func assertSameStringArray(expected, actual []string, t *testing.T) {
  if len(expected) != len(actual) {
    t.Fatalf("TestAlreadyFormattedString: Array lengths differ: expected %v, got %v.", len(expected), len(actual))
  }

  for i := 0; i < len(actual); i++ {
    if expected[i] != actual[i] {
      t.Errorf("TestAlreadyFormattedString: value in element %v is different. Expected '%v', got '%v'.", i, expected[i], actual[i])
    }
  }
}

func assertEncoded(expected []string, message string, t *testing.T) {
  actual := EncodeMessage(message)

  assertSameStringArray(expected, actual, t)
}
