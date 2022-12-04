package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type count struct {
	part1, part2 int
}

func fullyContained(r1, r2 [2]int) bool {
	if r1[0] >= r2[0] && r1[1] <= r2[1] {
		return true
	}
	if r2[0] >= r1[0] && r2[1] <= r1[1] {
		return true
	}

	return false
}

func overlap(r1, r2 [2]int) bool {
	if r1[1] <= r2[1] && r1[1] >= r2[0] {
		return true
	}
	if r2[1] <= r1[1] && r2[1] >= r1[0] {
		return true
	}

	return false
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s file.\n", os.Args[0])
		os.Exit(1)
	}

	input, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()
	count := count{part1: 0, part2: 0}

	for {
		var pair [2][2]int
		_, err := fmt.Fscanf(input, "%d-%d,%d-%d", &pair[0][0], &pair[0][1], &pair[1][0], &pair[1][1])

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}

			break
		}

		if fullyContained(pair[0], pair[1]) {
			count.part1++
		}
		if overlap(pair[0], pair[1]) {
			count.part2++
		}
	}

	fmt.Printf("Part 1: %d\n", count.part1)
	fmt.Printf("Part 2: %d\n", count.part2)
}
