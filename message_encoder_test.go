package main

import (
  "testing"
)

func assertSameArray(actual, expected []string, t *testing.T) {
  if len(actual) != len(expected) {
    t.Fatalf("TestAlreadyFormattedString: Array lengths differ: expected %v, got %v.", len(expected), len(actual))
  }

  for i := 0; i < len(actual); i++ {
    if actual[i] != expected[i] {
      t.Errorf("TestAlreadyFormattedString: value in element %v is different. Expected '%v', got '%v'.", i, expected[i], actual[i])
    }
  }
}

func TestAlreadyFormattedString(t *testing.T) {
  message := "HELLO WORLD"
  expected := []string{ "HELLO", "WORLD" }

  actual := EncodeMessage(message)

  assertSameArray(actual, expected, t)
}

func TestCovertsToUpperCase(t *testing.T) {
 message := "Hello World"
 expected := []string{ "HELLO", "WORLD" }

 actual := EncodeMessage(message)

 assertSameArray(actual, expected, t)
}

func TestRemovesNonAlphaCharacters(t *testing.T) {
  message := "Hello, World!"
  expected := []string{ "HELLO", "WORLD" }

  actual := EncodeMessage(message)

  assertSameArray(actual, expected, t)
}

func TestPadding(t *testing.T) {
  message := "Goodbye, cruel world!"
  expected := []string{ "GOODB", "YECRU", "ELWOR", "LDXXX" }

  actual := EncodeMessage(message)

  assertSameArray(actual, expected, t)
}