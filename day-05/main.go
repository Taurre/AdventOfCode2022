package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	part2 := flag.Bool("2", false, "Apply part 2 logic")
	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s file.\n", os.Args[0])
		os.Exit(1)
	}

	
	input, err := os.Open(flag.Args()[0])

	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()
	scanner := bufio.NewScanner(input)
	var crates [][]rune

	for scanner.Scan() {
		text := scanner.Text()

		if crates == nil {
			size := math.Ceil(float64(len(text)) / float64(4))
			crates = make([][]rune, int(size))
		}
		if text == "" {
			break
		}
		if strings.Index(scanner.Text(), "[") < 0 {
			continue
		}

		line := []byte(scanner.Text())
		stack := 0

		for i := 0; i < len(line); {
			if line[i] == '[' {
				crates[stack] = append([]rune{rune(line[i+1])}, crates[stack]...)
			}

			i += 4
			stack++
		}
	}

	if scanner.Err() != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		var count, from, to int
		_, err := fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &count, &from, &to)

		if err != nil {
			log.Fatal(err)
		}

		if *part2 {
			last := len(crates[from-1])
			crates[to-1] = append(crates[to-1], crates[from-1][last-count:last]...)
			crates[from-1] = crates[from-1][:last-count]
		} else {
			for i := 0; i < count; i++ {
				last := len(crates[from-1])-1
				crates[to-1] = append(crates[to-1], crates[from-1][last])
				crates[from-1] = crates[from-1][:last]
			}
		}
	}

	for _, crate := range crates {
		fmt.Printf("%c", crate[len(crate)-1])
	}

	fmt.Printf("\n")
}
