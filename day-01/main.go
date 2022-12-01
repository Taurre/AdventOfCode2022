package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func sum(s []int) (sum int) {
	for _, n := range s {
		sum += n
	}

	return sum
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

	sort.Ints(elves)
	fmt.Printf("Max of calories: %d\n", elves[len(elves)-1])
	fmt.Printf("Sum of top three: %d\n", sum(elves[len(elves)-3:]))
}
