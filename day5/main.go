package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack string

type Instruction struct {
  NumberOfCrates int
  Source int
  Destination int
}

type Input struct {
  Stacks []Stack
  Instructions []Instruction
}

func readInputFile(filename string) Input {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

  input := Input{Stacks: []Stack{}, Instructions: []Instruction{}}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
    trimmedLine := strings.Trim(line, " ")
    if strings.HasPrefix(trimmedLine, "[") { // parse stack content
      stackNumber := 0
      for i := 1 ; i < len(line); i += 4 {
        crateMarker := string(line[i])
        if crateMarker != " " {
          if len(input.Stacks) == 0 || len(input.Stacks) <= stackNumber {
            increaseSizeBy := stackNumber - len(input.Stacks) + 1
            for s := 0 ; s < increaseSizeBy; s++ {
              input.Stacks = append(input.Stacks, "")
            }
          }
          input.Stacks[stackNumber] = Stack(crateMarker + string(input.Stacks[stackNumber]))
        }
        stackNumber++
      }
    } else if strings.HasPrefix(line, "move") { // parse instruction
      parts := strings.Split(line, " ")
      numberOfCrates, _ := strconv.Atoi(parts[1])
      source, _ := strconv.Atoi(parts[3])
      destination, _ := strconv.Atoi(parts[5])
      instruction := Instruction{ NumberOfCrates: numberOfCrates, Source: source - 1, Destination: destination - 1 }

      input.Instructions = append(input.Instructions, instruction)
    }
  }
	return input
}

func partOne(filename string) {
  input := readInputFile(filename)

  // carry out instructions
  for _, instruction := range input.Instructions {
    for i := 0; i < instruction.NumberOfCrates; i++ {
      sourceStack := input.Stacks[instruction.Source]
      destinationStack := input.Stacks[instruction.Destination]
      crateMarker := sourceStack[len(sourceStack) - 1:]
      input.Stacks[instruction.Source] = sourceStack[:len(sourceStack) - 1]
      input.Stacks[instruction.Destination] = destinationStack + crateMarker
    }
  }

  result := ""
  for _, stack := range input.Stacks {
    result += string(stack[len(stack) - 1:])
  }
  fmt.Println(result)
}

func main() {
  partOne("input.txt")
}
