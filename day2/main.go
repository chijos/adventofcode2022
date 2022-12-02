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
const outcome_loss = "X"
const outcome_draw = "Y"
const outcome_win = "Z"
const win = 6
const draw = 3
const loss = 0
const rock = 1
const paper = 2
const scissors = 3

func GetScore(opponentShape string, outcome string) int {
	score := 0

	if outcome == outcome_loss {
		score += loss
	} else if outcome == outcome_win {
		score += win
	} else {
		score += draw
	}

	if opponentShape == opp_rock {
		if outcome == outcome_win {
			score += paper
		} else if outcome == outcome_loss {
			score += scissors
		} else {
			score += rock
		}
	} else if opponentShape == opp_paper {
		if outcome == outcome_win {
			score += scissors
		} else if outcome == outcome_loss {
			score += rock
		} else {
			score += paper
		}
	} else if opponentShape == opp_scissors {
		if outcome == outcome_win {
			score += rock
		} else if outcome == outcome_loss {
			score += paper
		} else {
			score += scissors
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
