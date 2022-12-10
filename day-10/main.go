package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type instruction struct {
	name    string
	operand int
}

func newInstruction(text string) instruction {
	if text == "noop" {
		return instruction{name: "noop"}
	}

	var instruction instruction
	_, err := fmt.Sscanf(text, "%s %d", &instruction.name, &instruction.operand)

	if err != nil {
		log.Fatalf("'%s': bad formated line.\n", text)
	}

	return instruction
}

func wipeScreen(screen []rune) {
	for i, _ := range screen {
		screen[i] = '.'
	}
}

func updateScreen(screen []rune, register, cycle int) {
	cycle = cycle % len(screen)

	if cycle >= register-1 && cycle <= register+1 {
		screen[cycle] = '#'
	}

}

func printScreen(screen []rune) {
	for _, c := range screen {
		fmt.Printf("%c", c)
	}

	fmt.Printf("\n")
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
	register := 1
	stops := []int{20, 60, 100, 140, 180, 220}
	cycle := 0
	sum := 0
	screen := make([]rune, 40)

	wipeScreen(screen)

	for scanner.Scan() {
		text := scanner.Text()
		instruction := newInstruction(text)
		cycle++
		updateScreen(screen, register, cycle-1)

		if cycle%len(screen) == 0 {
			printScreen(screen)
			wipeScreen(screen)
		}

		for _, stop := range stops {
			if cycle == stop {
				sum += cycle * register
			}
		}

		if instruction.name == "addx" {
			cycle++
			updateScreen(screen, register, cycle-1)

			if cycle%len(screen) == 0 {
				printScreen(screen)
				wipeScreen(screen)
			}

			for _, stop := range stops {
				if cycle == stop {
					sum += cycle * register
				}
			}

			register += instruction.operand
		}
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	fmt.Println(sum)
}
