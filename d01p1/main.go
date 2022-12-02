package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)

	var max, sum int
	for s.Scan() {
		t := s.Text()
		if t == "" {
			if sum > max {
				max = sum
			}
			sum = 0
		} else {
			i, err := strconv.Atoi(t)
			if err != nil {
				log.Fatal(err)
			}
			sum += i
		}
	}

	if s.Err() != nil {
		log.Fatal(s.Err())
	}

	if sum > max {
		max = sum
	}
	fmt.Println(max)
}
