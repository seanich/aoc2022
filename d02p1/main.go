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
	'X': MoveRock,
	'B': MovePaper,
	'Y': MovePaper,
	'C': MoveScissors,
	'Z': MoveScissors,
}

const (
	OutcomeLoss = iota * 3
	OutcomeDraw
	OutcomeWin
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)

	var score int
	for s.Scan() {
		runes := []rune(s.Text())
		theirMove, myMove := moveMap[runes[0]], moveMap[runes[2]]
		score += myMove

		if theirMove == myMove {
			score += OutcomeDraw
		} else if myMove == (theirMove)%3+1 {
			score += OutcomeWin
		} else {
			score += OutcomeLoss
		}
	}

	if s.Err() != nil {
		log.Fatal(s.Err())
	}

	fmt.Println(score)
}
