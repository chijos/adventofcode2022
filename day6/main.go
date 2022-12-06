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

func detectMarkerPacket(dataStream string, packetSize int) {
  for i := 0; i < len(dataStream) - packetSize; i++ {
    packet := dataStream[i:i+packetSize]
    if !hasRepeatedCharacters(packet) {
      fmt.Println(i + packetSize)
      return
    }
  }
}

func partOne(dataStream string) {
  detectMarkerPacket(dataStream, 4)
}

func partTwo(dataStream string) {
  detectMarkerPacket(dataStream, 14)
}

func main() {
	dataStream := readInputFile("input.txt")
	partOne(dataStream)
	partTwo(dataStream)
}
