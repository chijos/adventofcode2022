package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const priorityString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getPriority(item string) int {
  return strings.Index(priorityString, item) + 1
}

func getMisplacedItem(rucksack string) string {
  midpoint := len(rucksack) / 2
  for i := 0; i < midpoint; i++ {
    for j := midpoint; j < len(rucksack); j++ {
      if rucksack[i] == rucksack[j] {
        return string(rucksack[i])
      }
    }
  }
  return "" // should never reach here according to input specification
}

func readInputFile(filename string) []string {
  file, err := os.Open(filename)
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  var rucksacks []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    rucksacks = append(rucksacks, line)
  }
  return rucksacks
}

func PartOne(rucksacks []string) {
  sumOfPriorities := 0
  for _, rucksack := range rucksacks {
    misplacedItem := getMisplacedItem(rucksack)
    sumOfPriorities += getPriority(misplacedItem)
  }
  fmt.Println(sumOfPriorities)
}

func PartTwo(rucksacks []string) {
  sumOfPriorities := 0

  for i := 0; i < len(rucksacks); i+=3 {
    group := rucksacks[i:i+3]
    sort.Slice(group, func (i,j int) bool {
      return len(group[i]) > len(group[j])
    })
    for k := 0 ; k < len(group[0]); k++ {
      strItem := string(group[0][k])
      isInSecondRucksack := strings.Index(group[1], strItem) != -1
      isInThirdRucksack := strings.Index(group[2], strItem) != -1
      if isInSecondRucksack && isInThirdRucksack {
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
