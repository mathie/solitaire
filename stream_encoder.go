package main

func StreamEncoder(stream []string) [][]int {
  encodedValues := [][]int{}

  for i := 0; i < len(stream); i++ {
    block := stream[i]
    encodedBlock := []int{}
    for j := 0; j < len(block); j++ {
      char := block[j]
      encodedBlock = append(encodedBlock, int(char - 'A' + 1))
    }

    encodedValues = append(encodedValues, encodedBlock)
  }
  return encodedValues
}

func StreamDecoder(stream [][]int) []string {
  decodedValues := []string{}

  for i := 0; i < len(stream); i++ {
    block := stream[i]
    decodedBlock := []byte{}
    for j := 0; j < len(block); j++ {
      char := block[j]
      decodedBlock = append(decodedBlock, byte(char + 'A' - 1))
    }

    decodedValues = append(decodedValues, string(decodedBlock))
  }

  return decodedValues
}