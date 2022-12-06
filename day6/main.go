package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInputFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
  scanner.Scan()
  return scanner.Text()
}

func hasRepeatedCharacters(packet string) bool {
  for i := 0; i < len(packet) - 1; i++ {
    character := string(packet[i])
    for j := i + 1; j < len(packet); j++ {
      otherCharacter := string(packet[j])
      if otherCharacter == character {
        return true
      }
    }
  }
  return false
}

func partOne(dataStream string) {
  for i := 0; i < len(dataStream) - 4; i++ {
    packet := dataStream[i:i+4]
    if !hasRepeatedCharacters(packet) {
      fmt.Println(i + 4)
      return
    }
  }
}

func partTwo(dataStream string) {
  for i := 0; i < len(dataStream) - 14; i++ {
    packet := dataStream[i:i+14]
    if !hasRepeatedCharacters(packet) {
      fmt.Println(i + 14)
      return
    }
  }
}

func main() {
	dataStream := readInputFile("input.txt")
	partOne(dataStream)
	partTwo(dataStream)
}
