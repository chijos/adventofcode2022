package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Assignment struct {
	Start int
	End   int
}

type AssignmentPair struct {
	A Assignment
	B Assignment
}

func (a AssignmentPair) HasRedundantAssignment() bool {
  if a.A.Start <= a.B.Start && a.A.End >= a.B.End {
    return true
  } else if a.B.Start <= a.A.Start && a.B.End >= a.A.End {
    return true
  } else {
    return false
  }
}

func parseAssignment(input string) Assignment {
	split := strings.Split(input, "-")
	start, _ := strconv.Atoi(split[0])
	end, _ := strconv.Atoi(split[1])
	return Assignment{Start: start, End: end}
}

func parseAssignmentPair(input string) AssignmentPair {
  split := strings.Split(input, ",")
  return AssignmentPair{A: parseAssignment(split[0]), B: parseAssignment(split[1])}
}

func readInputFile(filename string) []AssignmentPair {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var assignmentPairs []AssignmentPair
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
    assignmentPair := parseAssignmentPair(line)
		assignmentPairs = append(assignmentPairs, assignmentPair)
	}
	return assignmentPairs
}

func partOne(fileName string) {
  assignmentPairs := readInputFile(fileName)
  numberOfRedundantAssignments := 0
  for _, pair := range assignmentPairs {
    if pair.HasRedundantAssignment() {
      numberOfRedundantAssignments += 1
    }
  }
  fmt.Println(numberOfRedundantAssignments)
}

func partTwo(fileName string) {

}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
