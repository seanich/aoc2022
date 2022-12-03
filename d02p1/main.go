package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type move int

const (
	MoveRock move = iota + 1
	MovePaper
	MoveScissors
)

var moveMap = map[rune]move{
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
		score += int(myMove)
		if theirMove == myMove {
			score += OutcomeDraw
		} else if myMove == MoveRock && theirMove == MoveScissors || myMove == MovePaper && theirMove == MoveRock || myMove == MoveScissors && theirMove == MovePaper {
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
