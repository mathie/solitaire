package main

import "testing"

func TestEncryption(t *testing.T) {
  message   := []string{ "CODEI", "NRUBY", "LIVEL", "ONGER" }
  expected  := []string{ "GLNCQ", "MJAFF", "FVOMB", "JIYCB" }

  deck := NewUnkeyedDeck()
  actual := Encrypt(message, deck)

  assertSameStringArray(expected, actual, t)
}

func TestDecryption(t *testing.T) {
  message   := []string{ "GLNCQ", "MJAFF", "FVOMB", "JIYCB" }
  expected  := []string{ "CODEI", "NRUBY", "LIVEL", "ONGER" }

  deck := NewUnkeyedDeck()
  actual := Decrypt(message, deck)

  assertSameStringArray(expected, actual, t)
}

func ExampleHelloWorld() {
  main()

  // Output:
  // Hello World.
}