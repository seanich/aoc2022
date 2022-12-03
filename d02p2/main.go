package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	MoveRock = iota + 1
	MovePaper
	MoveScissors
)

var moveMap = map[rune]int{
	'A': MoveRock,
	'B': MovePaper,
	'C': MoveScissors,
}

const (
	OutcomeLoss = iota * 3
	OutcomeDraw
	OutcomeWin
)

var outcomeMap = map[rune]int{
	'X': OutcomeLoss,
	'Y': OutcomeDraw,
	'Z': OutcomeWin,
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)

	var score int
	for s.Scan() {
		runes := []rune(s.Text())
		theirMove, myOutcome := moveMap[runes[0]], outcomeMap[runes[2]]
		score += myOutcome
		switch myOutcome {
		case OutcomeDraw:
			score += theirMove
		case OutcomeLoss:
			score += (theirMove+1)%3 + 1
		case OutcomeWin:
			score += theirMove%3 + 1
		}
	}

	if s.Err() != nil {
		log.Fatal(s.Err())
	}

	fmt.Println(score)
}
