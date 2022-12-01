package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func max(s []int) (offset, max int) {
	for i, n := range s {
		if max < n {
			max = n
			offset = i
		}
	}

	return offset, max
}

func topThree(s []int) (one, two, three int) {
	var offset int

	offset, one = max(s)
	s = append(s[:offset], s[offset+1:]...)
	offset, two = max(s)
	s = append(s[:offset], s[offset+1:]...)
	offset, three = max(s)
	return one, two, three
}

func main() {
	input, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)
	elves := make([]int, 1)

	for current := 0; scanner.Scan(); {
		if scanner.Text() == "" {
			elves = append(elves, 0)
			current++
			continue
		}

		calories, err := strconv.ParseUint(scanner.Text(), 10, 64)

		if err != nil {
			log.Fatal(err)
		}

		elves[current] += int(calories)
	}

	_, max := max(elves)
	fmt.Printf("Max of calories: %d\n", max)
	one, two, three := topThree(elves)
	fmt.Printf("Sum of top three: %d\n", one+two+three)
}
