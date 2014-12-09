package main

import (
  "math"
  "bufio"
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
func EncodeMessage(message string) []string {
  removeNonAlpha := func(r rune) rune {
    if r >= 'A' && r <= 'Z'{
      return r
    } else {
      return -1
    }
  }

  sanitizedString := strings.Map(removeNonAlpha, strings.ToUpper(message))

  scanner := bufio.NewScanner(strings.NewReader(sanitizedString))

  split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
    if len(data) > 5 {
      return 5, data[0:5], nil
    } else if atEOF {
      padding := []byte{'X','X','X','X','X'}
      return len(data), append(data[:], padding[:(5-len(data))]...), nil
    } else {
      return 0, nil, nil
    }
  }

  scanner.Split(split)

  i := 0
  results := make([]string, int32(math.Ceil(float64(len(sanitizedString)) / 5.0)))

  for scanner.Scan() {
    results[i] = scanner.Text()
    i++
  }

  return results
}
