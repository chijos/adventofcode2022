package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const opp_rock = "A"
const opp_paper = "B"
const opp_scissors = "C"
const rock = "X"
const paper = "Y"
const scissors = "Z"
const win = 6
const draw = 3
const loss = 0

func GetScore(opponentShape string, myShape string) int {
	score := 0

	if myShape == rock {
		score += 1
	} else if myShape == paper {
		score += 2
	} else if myShape == scissors {
		score += 3
	}

	if myShape == rock {
		if opponentShape == opp_rock {
			score += draw
		} else if opponentShape == opp_scissors {
			score += win
		} else {
			score += loss
		}
	} else if myShape == paper {
		if opponentShape == opp_rock {
			score += win
		} else if opponentShape == opp_paper {
			score += draw
		} else {
			score += loss
		}
	} else if myShape == scissors {
		if opponentShape == opp_rock {
			score += loss
		} else if opponentShape == opp_paper {
			score += win
		} else {
			score += draw
		}
	}

	return score
}

type Play struct {
	OpponentShape string
	MyShape       string
}

func ReadStrategyFile(filename string) []Play {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var plays []Play
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		shapes := strings.Split(line, " ")
		plays = append(plays, Play{shapes[0], shapes[1]})
	}

	return plays
}

func main() {
	plays := ReadStrategyFile("input.txt")

	totalScore := 0
	for i := 0; i < len(plays); i++ {
		score := GetScore(plays[i].OpponentShape, plays[i].MyShape)
		totalScore += score
	}
	fmt.Println(totalScore)
}
