package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	scorePart1 = map[string]map[string]int{
		"X": {
			"A": 4,
			"B": 1,
			"C": 7,
		},
		"Y": {
			"A": 8,
			"B": 5,
			"C": 2,
		},
		"Z": {
			"A": 3,
			"B": 9,
			"C": 6,
		},
	}
	scorePart2 = map[string]map[string]int{
		"X": {
			"A": 3,
			"B": 1,
			"C": 2,
		},
		"Y": {
			"A": 4,
			"B": 5,
			"C": 6,
		},
		"Z": {
			"A": 8,
			"B": 9,
			"C": 7,
		},
	}
)

func readRound(scanner *bufio.Scanner) (rune, rune, error) {
	shapes := []rune{0, 0}

	for i := 0; i < 2; i++ {
		if !scanner.Scan() {
			if scanner.Err() == nil {
				return 0, 0, io.EOF
			} else {
				return 0, 0, scanner.Err()
			}
		}

		shapes[i] = rune(scanner.Text()[0])
	}

	return shapes[0], shapes[1], nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s file\n", os.Args[0])
		os.Exit(1)
	}

	input, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	total := make([]int, 2)

	for {
		left, right, err := readRound(scanner)

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}

			break
		}

		total[0] += scorePart1[string(right)][string(left)]
		total[1] += scorePart2[string(right)][string(left)]
	}

	fmt.Printf("Score part 1: %d\n", total[0])
	fmt.Printf("Score part 2: %d\n", total[1])
}
