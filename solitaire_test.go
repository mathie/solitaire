package main

import "testing"

func TestEncryption(t *testing.T) {
  message   := []string{ "CODEI", "NRUBY", "LIVEL", "ONGER" }
  keystream := []string{ "DWJXH", "YRFDG", "TMSHP", "UURXJ" }
  expected  := []string{ "GLNCQ", "MJAFF", "FVOMB", "JIYCB" }

  actual := Encrypt(message, keystream)

  assertSameStringArray(expected, actual, t)
}

func TestDecryption(t *testing.T) {
  message   := []string{ "GLNCQ", "MJAFF", "FVOMB", "JIYCB" }
  keystream := []string{ "DWJXH", "YRFDG", "TMSHP", "UURXJ" }
  expected  := []string{ "CODEI", "NRUBY", "LIVEL", "ONGER" }

  actual := Decrypt(message, keystream)

  assertSameStringArray(expected, actual, t)
}

func ExampleHelloWorld() {
  main()

  // Output:
  // Hello World.
}