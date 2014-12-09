package main

import "testing"

func TestStreamEncoding(t *testing.T) {
  message := []string{ "CODEI", "NRUBY", "LIVEL", "ONGER" }
  expected := [][]int{ {3, 15, 4, 5, 9}, {14, 18, 21, 2, 25}, {12, 9, 22, 5, 12}, {15, 14, 7, 5, 18} }

  actual := StreamEncoder(message)

  assertSameIntArray(expected, actual, t)
}

func TestStreamDecoding(t *testing.T) {
  message := [][]int{ {3, 15, 4, 5, 9}, {14, 18, 21, 2, 25}, {12, 9, 22, 5, 12}, {15, 14, 7, 5, 18} }
  expected := []string{ "CODEI", "NRUBY", "LIVEL", "ONGER" }

  actual := StreamDecoder(message)

  assertSameStringArray(expected, actual, t)
}

func assertSameIntArray(expected, actual [][]int, t *testing.T) {
  if len(expected) != len(actual) {
    t.Fatalf("TestAlreadyFormattedString: Array lengths differ: expected %v, got %v.", len(expected), len(actual))
  }

  for i := 0; i < len(actual); i++ {
    expectedBlock := expected[i]
    actualBlock := actual[i]

    if len(expectedBlock) != len(actualBlock) {
      t.Fatalf("TestAlreadyFormattedString: Array lengths differ for block %v: expected %v, got %v.", i, len(expectedBlock), len(actualBlock))
    }

    for j := 0; j < len(actualBlock); j++ {
      if expectedBlock[j] != actualBlock[j] {
        t.Errorf("TestAlreadyFormattedString: value in element %v is different. Expected '%v', got '%v'.", i, expectedBlock[j], actualBlock[j])
      }
    }
  }
}
