package main

import (
  "testing"
)


func TestAlreadyFormattedString(t *testing.T) {
 message := "HELLO WORLD"
 expected := "HELLOWORLD"

 actual := EncodeMessage(message)

 if actual != expected {
   t.Fatalf("TestAlreadyFormattedString() failed. Expected '%v', got '%v'.", expected, actual)
 }
}

func TestCovertsToUpperCase(t *testing.T) {
 message := "Hello World"
 expected := "HELLOWORLD"

 actual := EncodeMessage(message)

 if actual != expected {
   t.Fatalf("TestCovertsToUpperCase() failed. Expected '%v', got '%v'.", expected, actual)
 }
}

func TestRemovesNonAlphaCharacters(t *testing.T) {
  message := "Hello, World!"
  expected := "HELLOWORLD"

  actual := EncodeMessage(message)

  if actual != expected {
    t.Fatalf("TestRemovesNonAlphaCharacters() failed. Expected '%v', got '%v'.", expected, actual)
  }
}