package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println("Hello World.")
}

func OutputEncryptedString(message string) {
  encryptedMessage := Encrypt(EncodeMessage(message), NewUnkeyedDeck())
  fmt.Println(strings.Join(encryptedMessage, " "))
}

func OutputDecryptedString(message string) {
  decryptedMessage := Decrypt(EncodeMessage(message), NewUnkeyedDeck())
  fmt.Println(strings.Join(decryptedMessage, " "))
}

func Encrypt(message []string, deck *Deck) []string {
  encryptor := func(messageChar, keystreamChar int) int {
    return (messageChar + keystreamChar) % 26
  }

  return process(message, deck, encryptor)
}

func Decrypt(message []string, deck *Deck) []string {
  decryptor := func(messageChar, keystreamChar int) int {
    result := messageChar - keystreamChar
    if result <= 0 {
      result += 26
    }
    return result
  }

  return process(message, deck, decryptor)
}

type processFunc func(int, int) int

func process(message []string, deck *Deck, processFunc processFunc) []string {
  encodedMessage := StreamEncoder(message)

  results := [][]int{}

  for i := 0; i < len(encodedMessage); i++ {
    messageBlock := encodedMessage[i]
    resultsBlock := []int{}

    for j := 0; j < len(messageBlock); j++ {
      keystreamChar := int(deck.GetNextCode() % 26)
      resultsBlock = append(resultsBlock, processFunc(messageBlock[j], keystreamChar))
    }

    results = append(results, resultsBlock)
  }

  return StreamDecoder(results)
}
