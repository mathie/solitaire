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

func (deck *Deck) MoveAJokerDown() {
  deck.moveCardDown(jokerA, 1)
}

func (deck *Deck) MoveBJokerDown() {
  deck.moveCardDown(jokerB, 2)
}

func (deck *Deck) TripleCut() {
  firstIndex := bytes.IndexByte(deck.cards, jokerA)
  secondIndex := bytes.IndexByte(deck.cards, jokerB)

  if firstIndex > secondIndex {
    firstIndex, secondIndex = secondIndex, firstIndex
  }

  minSlice := deck.cards[0:firstIndex]
  midSlice := deck.cards[firstIndex:secondIndex + 1]
  maxSlice := deck.cards[secondIndex + 1:]

  deck.cards = append(append(maxSlice, midSlice...), minSlice...)
}

func (deck *Deck) CountCut() {
  count := deck.cards[len(deck.cards) - 1]

  minSlice := deck.cards[0:count]
  midSlice := deck.cards[count:len(deck.cards) - 1]

  deck.cards = append(append(midSlice, minSlice...), count)
}

func (deck Deck) GetCode() byte {
  position := deck.cards[0]

  // Only count 53 cards down for either joker, since you can't count 54 cards
  // down from here in a 54 card stack!
  if position == jokerB {
    return deck.cards[position - 1]
  } else {
    return deck.cards[position]
  }
}

func (deck *Deck) GetNextCode() byte {
  deck.MoveAJokerDown()
  deck.MoveBJokerDown()
  deck.TripleCut()
  deck.CountCut()

  result := deck.GetCode()
  if result >= jokerA {
    return deck.GetNextCode()
  } else {
    return result
  }
}

func (deck *Deck) moveCardDown(card byte, distance int) {
  for i := 0; i < distance; i++ {
    jokerIndex := bytes.IndexByte(deck.cards, card)
    if jokerIndex == len(deck.cards) - 1 {
      deck.cards = append(deck.cards[len(deck.cards) - 1:], deck.cards[:len(deck.cards) - 1]...)

      // Need to move once more, since we've really just moved the card to the
      // top of the stack.
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