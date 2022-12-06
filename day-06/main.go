package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func isMarker(data string) bool {
	for _, r := range data {
		if strings.Count(data, string(r)) > 1 {
			return false
		}
	}	

	return true
}

func main() {
	log.SetFlags(0)

	if len(os.Args) == 1 {
		log.Fatalf("Usage: %s file.\n", os.Args[0])
	}

	input, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		text := scanner.Text()

		for count := 4; count < len(text); count++ {
			current := text[count-4:count]

			if isMarker(current) {
				fmt.Printf("'%s' is a start-of-packet marker starting at %d\n", current, count)
				break
			}
		}

		for count := 14; count < len(text); count++ {
			current := text[count-14:count]

			if isMarker(current) {
				fmt.Printf("'%s' is a start-of-message marker starting at %d\n", current, count)
				break
			}
		}
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
