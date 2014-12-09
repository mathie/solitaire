package main

import (
  "fmt"
  "log"
)

func Encrypt(message, keystream []string) []string {
  encodedMessage := StreamEncoder(message)
  encodedKeystream := StreamEncoder(keystream)

  results := [][]int{}

  if len(encodedMessage) != len(encodedKeystream) {
    log.Panicf("Message and key stream are different lengths.")
  }

  for i := 0; i < len(encodedMessage); i++ {
    messageBlock := encodedMessage[i]
    keystreamBlock := encodedKeystream[i]
    resultsBlock := []int{}

    if len(messageBlock) != len(keystreamBlock) {
      log.Panicf("Message block and key stream blocks are different lengths.")
    }

    for j := 0; j < len(messageBlock); j++ {
      messageChar   := messageBlock[j]
      keystreamChar := keystreamBlock[j]
      result := (messageChar + keystreamChar) % 26
      resultsBlock = append(resultsBlock, result)
    }

    results = append(results, resultsBlock)
  }

  return StreamDecoder(results)
}

func Decrypt(message, keystream []string) []string {
  encodedMessage := StreamEncoder(message)
  encodedKeystream := StreamEncoder(keystream)

  results := [][]int{}

  if len(encodedMessage) != len(encodedKeystream) {
    log.Panicf("Message and key stream are different lengths.")
  }

  for i := 0; i < len(encodedMessage); i++ {
    messageBlock := encodedMessage[i]
    keystreamBlock := encodedKeystream[i]
    resultsBlock := []int{}

    if len(messageBlock) != len(keystreamBlock) {
      log.Panicf("Message block and key stream blocks are different lengths.")
    }

    for j := 0; j < len(messageBlock); j++ {
      messageChar   := messageBlock[j]
      keystreamChar := keystreamBlock[j]
      result := messageChar - keystreamChar
      if result < 0 {
        result += 26
      }
      resultsBlock = append(resultsBlock, result)
    }

    results = append(results, resultsBlock)
  }

  return StreamDecoder(results)
}

func main() {
  fmt.Println("Hello World.")
}