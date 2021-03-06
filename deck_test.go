package main

import (
  "testing"
)

func TestInitialisationUnkeyedDeck(t *testing.T) {
  deck := NewUnkeyedDeck()
  if len(deck.cards) != 54 {
    t.Fatalf("Unkeyed deck was initialised with the wrong number of cards.")
  }
}

func TestMoveAJokerDown(t *testing.T) {
  deck := NewUnkeyedDeck()
  expected := []byte{
     1,  2,  3,  4,  5,  6,  7,  8,  9, 10, 11, 12, 13,
    14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
    27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
    40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52,
    54, 53,
  }
  deck.MoveAJokerDown()

  assertCardsEqual(expected, deck.cards, t)
}

func TestThenMoveBJokerDown(t *testing.T) {
  deck := NewUnkeyedDeck()
  expected := []byte{
     1, 54,  2,  3,  4,  5,  6,  7,  8,  9, 10, 11, 12,
    13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
    26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38,
    39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
    52, 53,
  }
  deck.MoveAJokerDown()
  deck.MoveBJokerDown()

  assertCardsEqual(expected, deck.cards, t)
}

func TestThenTripleCut(t *testing.T) {
  deck := NewUnkeyedDeck()
  expected := []byte{
    54,  2,  3,  4,  5,  6,  7,  8,  9, 10, 11, 12, 13,
    14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
    27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
    40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52,
    53,  1,
  }
  deck.MoveAJokerDown()
  deck.MoveBJokerDown()
  deck.TripleCut()

  assertCardsEqual(expected, deck.cards, t)
}

func TestThenCountCut(t *testing.T) {
  deck := NewUnkeyedDeck()
  expected := []byte{
     2,  3,  4,  5,  6,  7,  8,  9, 10, 11, 12, 13, 14,
    15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27,
    28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
    41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53,
    54,  1,
  }
  deck.MoveAJokerDown()
  deck.MoveBJokerDown()
  deck.TripleCut()
  deck.CountCut()

  assertCardsEqual(expected, deck.cards, t)
}

func TestThenGetCode(t *testing.T) {
  deck := NewUnkeyedDeck()
  deck.MoveAJokerDown()
  deck.MoveBJokerDown()
  deck.TripleCut()
  deck.CountCut()

  actual := deck.GetCode()
  if actual != 4 {
    t.Fatalf("Expected code 4, got code %v", actual)
  }
}

func TestGetNextCode(t *testing.T) {
  deck := NewUnkeyedDeck()

  assertCodeEqual := func(expected, actual byte) {
    if actual != expected {
      t.Errorf("Expected code %v, got code %v", expected, actual)
    }
  }

  // First 10 codes from an unkeyed deck.
  assertCodeEqual( 4, deck.GetNextCode())
  assertCodeEqual(49, deck.GetNextCode())
  assertCodeEqual(10, deck.GetNextCode())
  assertCodeEqual(24, deck.GetNextCode())
  assertCodeEqual( 8, deck.GetNextCode())
  assertCodeEqual(51, deck.GetNextCode())
  assertCodeEqual(44, deck.GetNextCode())
  assertCodeEqual( 6, deck.GetNextCode())
  assertCodeEqual( 4, deck.GetNextCode())
  assertCodeEqual(33, deck.GetNextCode())
}

func TestGetNextChar(t *testing.T) {
  deck := NewUnkeyedDeck()

  assertCharEqual := func(expected, actual byte) {
    if actual != expected {
      t.Errorf("Expected %c, got %c", expected, actual)
    }
  }

  assertCharEqual('D', deck.GetNextChar())
  assertCharEqual('W', deck.GetNextChar())
  assertCharEqual('J', deck.GetNextChar())
  assertCharEqual('X', deck.GetNextChar())
  assertCharEqual('H', deck.GetNextChar())
  assertCharEqual('Y', deck.GetNextChar())
  assertCharEqual('R', deck.GetNextChar())
  assertCharEqual('F', deck.GetNextChar())
  assertCharEqual('D', deck.GetNextChar())
  assertCharEqual('G', deck.GetNextChar())
}

func assertCardsEqual(expected, actual []byte, t *testing.T) {
  const numberOfCards = 54

  if len(expected) != numberOfCards {
    t.Fatalf("Expected deck should have %v cards, only had %v", numberOfCards, len(expected))
  }

  if len(actual) != numberOfCards {
    t.Fatalf("Actual deck should have %v cards, only had %v", numberOfCards, len(actual))
  }

  for i := 0; i < numberOfCards; i++ {
    if expected[i] != actual[i] {
      t.Errorf("Expected card %v to be %v, but it was %v", i, expected[i], actual[i])
    }
  }
}