package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)

	var sum int
	var sums []int
	for s.Scan() {
		t := s.Text()
		if t == "" {
			sums = append(sums, sum)
			sum = 0
		} else {
			i, err := strconv.Atoi(t)
			if err != nil {
				log.Fatal(err)
			}
			sum += i
		}
	}
	sums = append(sums, sum)

	if s.Err() != nil {
		log.Fatal(s.Err())
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	sum = 0
	for _, cals := range sums[:3] {
		sum += cals
	}

	fmt.Println(sum)
}
