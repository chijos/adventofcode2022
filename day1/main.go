package main

import (
	"fmt"
  "os"
  "bufio"
  "strconv"
  "sort"
)

func ReadInputFile(filename string) []int {
  file, err := os.Open(filename)
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  var elves []int
  scanner := bufio.NewScanner(file)
  currentSum := 0
  for scanner.Scan() {
    line := scanner.Text()
    if line == "" {
      elves = append(elves, currentSum)
      currentSum = 0
    } else {
      calories, _ := strconv.Atoi(line)
      currentSum += calories
    }
  }
  if(currentSum > 0) {
    elves = append(elves, currentSum)
  }

  return elves
}

func PartOne(elves []int) {
  fmt.Println("== Part One ==")
  fmt.Println(elves[0])
}

func PartTwo(elves []int) {
  fmt.Println("== Part Two ==")
  sum := 0
  for _, c := range elves[:3] {
    sum += c
  }
  fmt.Println(sum)
}

func main() {
  elves := ReadInputFile("input.txt")
  sort.Sort(sort.Reverse(sort.IntSlice(elves[:])))
  PartOne(elves)
  PartTwo(elves)
}
