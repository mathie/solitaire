package main

import (
  "fmt"
  "log"
)

func main() {
  fmt.Println("Hello World.")
}

func Encrypt(message, keystream []string) []string {
  encryptor := func(messageChar, keystreamChar int) int {
    return (messageChar + keystreamChar) % 26
  }

  return process(message, keystream, encryptor)
}

func Decrypt(message, keystream []string) []string {
  decryptor := func(messageChar, keystreamChar int) int {
    result := messageChar - keystreamChar
    if result < 0 {
      result += 26
    }
    return result
  }

  return process(message, keystream, decryptor)
}

type processFunc func(int, int) int

func process(message, keystream []string, processFunc processFunc) []string {
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
      resultsBlock = append(resultsBlock, processFunc(messageBlock[j], keystreamBlock[j]))
    }

    results = append(results, resultsBlock)
  }

  return StreamDecoder(results)
}
