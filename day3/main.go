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

func (r Rucksack) Length() int {
  return len(r.CompartmentOne) + len(r.CompartmentTwo)
}

func (r Rucksack) AllItems() string {
  return r.CompartmentOne + r.CompartmentTwo
}

const priorityString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getPriority(item string) int {
  return strings.Index(priorityString, item) + 1
}

func getMisplacedItem(rucksack Rucksack) string {
  for i := 0; i < len(rucksack.CompartmentOne); i++ {
    item := string(rucksack.CompartmentOne[i])
    if strings.Index(rucksack.CompartmentTwo, item) > -1 {
      return item
    }
  }
  return "0" // should never get here according to data specifications
}

func readInputFile(filename string) []Rucksack {
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

func getRucksackWithMostItems(rucksacks []Rucksack) Rucksack {
  maxLen := -1
  var rucksackWithMostItems Rucksack
  for _, rucksack := range rucksacks {
    if rucksack.Length() > maxLen {
      maxLen = rucksack.Length()
      rucksackWithMostItems = rucksack
    }
  }
  return rucksackWithMostItems
}

func PartOne(rucksacks []Rucksack) {
  sumOfPriorities := 0
  for _, rucksack := range rucksacks {
    misplacedItem := getMisplacedItem(rucksack)
    sumOfPriorities += getPriority(misplacedItem)
  }
  fmt.Println(sumOfPriorities)
}
  
func PartTwo(rucksacks []Rucksack) {
  sumOfPriorities := 0

  for i := 0 ; i < len(rucksacks); i+=3 {
    group := rucksacks[i:i+3]
    rucksackWithMostItems := getRucksackWithMostItems(group)

    for _, item := range rucksackWithMostItems.AllItems() {
      strItem := string(item)
      isInFirstRucksack := strings.Index(group[0].AllItems(), strItem) != -1
      isInSecondRucksack := strings.Index(group[1].AllItems(), strItem) != -1
      isInThirdRucksack := strings.Index(group[2].AllItems(), strItem) != -1
      if isInFirstRucksack && isInSecondRucksack && isInThirdRucksack {
        sumOfPriorities += getPriority(strItem)
        break
      }
    }
  }
  
  fmt.Println(sumOfPriorities)
}

func main() {
  rucksacks := readInputFile("input.txt")

  PartOne(rucksacks)
  PartTwo(rucksacks)
}
