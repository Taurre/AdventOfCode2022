package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type sum struct {
	part1, part2 int
}

var (
	priority = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}
)

func findCommon(group []string) rune {
	count := make(map[string]int)
	offset := 0
	common := 0

	for _, s := range group {
		for _, item := range s {
			count[string(item)] |= 1 << offset
		}

		common |= 1 << offset
		offset++
	}

	for item, value := range count {
		if value == common {
			return rune(item[0])
		}
	}

	log.Fatalf("Cannot find common item for '%v'\n", group)
	return 0
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
	group := make([]string, 0)
	sum := sum{part1: 0, part2: 0}

	for scanner.Scan() {
		content := scanner.Text()
		first := content[:(len(content) / 2)]
		second := content[(len(content) / 2):]
		common := findCommon([]string{first, second})
		fmt.Println(string(common))
		sum.part1 += priority[string(common)]
		group = append(group, content)

		if len(group) == 3 {
			common = findCommon(group)
			sum.part2 += priority[string(common)]
			group = group[:0]
		}
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	fmt.Printf("Sum part 1: %d\n", sum.part1)
	fmt.Printf("Sum part 2: %d\n", sum.part2)
}
