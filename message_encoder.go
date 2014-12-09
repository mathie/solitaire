package main

import (
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
  blockSize := 5

  return encodeToBlocks(strings.Map(removeNonAlpha, strings.ToUpper(message)), blockSize)
}

// If the character is an upper case letter, return it. If not, return an
// invalid rune, indicating that it should be removed. Intended to be used with
// strings.Map().
func removeNonAlpha(r rune) rune {
  if r >= 'A' && r <= 'Z'{
    return r
  } else {
    return -1
  }
}

// Take a string and encode it into blocks, each of size blockSize. The last
// block will be padded with 'X' characters to make sure all blocks are of the
// same size.
func encodeToBlocks(message string, blockSize int) (results []string) {
  scanner := bufio.NewScanner(strings.NewReader(message))
  scanner.Split(splitIntoBlocks(blockSize))

  for scanner.Scan() {
    results = append(results, scanner.Text())
  }

  return results
}

// Split a string into blocks of blockSize bytes each. Returns a function which
// is intended to be used as the split function of a Scanner.
func splitIntoBlocks(blockSize int) bufio.SplitFunc {
  splitIntoBlocks := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
    err = nil

    if len(data) > blockSize {
      advance = blockSize
      token = data[0:5]
      return
    } else if atEOF {
      advance = len(data)
      token = padBlock(data, blockSize)
      return
    } else {
      // Not at the EOF, but we've been supplied less data than the block size.
      // Ask the scanner to try again with more data.
      return 0, nil, nil
    }
  }

  return splitIntoBlocks
}

// Take a block which is smaller than blockSize, and pad it up to being
// blockSize characters by appending 'X'es.
func padBlock(data []byte, blockSize int) []byte {
  // FIXME: This needs to be at least as large as the largest passed-in block
  // size. It just happens to work for this particular version...
  paddingBytes := []byte{'X','X','X','X','X'}

  paddingSize := blockSize - len(data)
  padding := paddingBytes[:paddingSize]

  return append(data, padding...)
}
