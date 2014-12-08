package main

import (
  "strings"
)

// The message encoder is responsible for taking an arbitrary string message,
// and turning it into something that the Solitaire Cipher can correctly encode.
// This process involves:
//
// * Converting the string to upper case.
// * Removing all characters that aren't uppercase letters.
// * Split the message into groups of five characters each, padding the last
//   group, if necessary.
//
// So, for example, given the following input string:
//
//  Coding in Go is awesome!
//
// it will produce the result:
//
// CODIN GINGO ISAWE SOMEX
func EncodeMessage(message string) string {
  removeNonAlpha := func(r rune) rune {
    if r >= 'A' && r <= 'Z'{
      return r
    } else {
      return -1
    }
  }

  return strings.Map(removeNonAlpha, strings.ToUpper(message))
}