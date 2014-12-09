package main

import (
  "bytes"
)

type Deck struct {
  cards []byte
}

// The two jokers are represented by the numbers 53 & 54; the rest of the cards
// are sequentially numbered as if they were in ascending order by suit (clubs,
// diamonds, hearts, then spades).
const jokerA = 53
const jokerB = 54

// Return a new, unkeyed deck -- ie all the cards are still sorted in ascending
// order by suit, followed by the A, then the B joker.
func NewUnkeyedDeck() *Deck {
  deck := &Deck{}

  for i := 1; i <= 54; i++ {
    deck.cards = append(deck.cards, byte(i))
  }

  return deck
}

func (deck Deck) MoveAJokerDown() {
  distance := 1
  jokerIndex := bytes.IndexByte(deck.cards, jokerA)
  movingCard := deck.cards[jokerIndex]
  nextCardIndex := jokerIndex + distance
  nextCard := deck.cards[nextCardIndex]
  deck.cards[jokerIndex] = nextCard
  deck.cards[nextCardIndex] = movingCard
}

func (deck *Deck) MoveBJokerDown() {
  distance := 2
  for i := 0; i < distance; i++ {
    jokerIndex := bytes.IndexByte(deck.cards, jokerB)
    if jokerIndex == len(deck.cards) - 1 {
      deck.cards = append(deck.cards[len(deck.cards) - 1:], deck.cards[:len(deck.cards) - 1]...)
      distance += 1
    } else {
      movingCard := deck.cards[jokerIndex]
      nextCardIndex := jokerIndex + 1
      nextCard := deck.cards[nextCardIndex]
      deck.cards[jokerIndex] = nextCard
      deck.cards[nextCardIndex] = movingCard
    }
  }
}