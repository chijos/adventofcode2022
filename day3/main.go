package main

import (
	"bufio"
	"fmt"
	"os"
  "strings"
)

type Rucksack struct {
  CompartmentOne string
  CompartmentTwo string
}

const priorityString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GetPriority(item string) int {
  return strings.Index(priorityString, item) + 1
}

func GetMisplacedItem(rucksack Rucksack) string {
  for i := 0; i < len(rucksack.CompartmentOne); i++ {
    item := string(rucksack.CompartmentOne[i])
    if strings.Index(rucksack.CompartmentTwo, item) > -1 {
      return item
    }
  }
  return "0" // should never get here according to data specifications
}

func ReadInputFile(filename string) []Rucksack {
  file, err := os.Open(filename)
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  var rucksacks []Rucksack
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    itemCount := len(line)
    itemPerCompartment := itemCount / 2
    compartmentOne := line[:itemPerCompartment]
    compartmentTwo := line[itemPerCompartment:]

    rucksacks = append(rucksacks, Rucksack{ CompartmentOne: compartmentOne, CompartmentTwo: compartmentTwo})
  }
  return rucksacks
}

func PartOne() {
  
}

func main() {
  rucksacks := ReadInputFile("input.txt")
  sumOfPriorities := 0
  for _, rucksack := range rucksacks {
    misplacedItem := GetMisplacedItem(rucksack)
    sumOfPriorities += GetPriority(misplacedItem)
  }
  fmt.Println(sumOfPriorities)
}
