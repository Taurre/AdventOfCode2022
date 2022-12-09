package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type knot struct {
	x, y int
}

func isAdjacent(head, tail *knot) bool {
	for _, x := range []int{-1, 0, 1} {
		for _, y := range []int{-1, 0, 1} {
			if tail.x+x == head.x && tail.y+y == head.y {
				return true
			}
		}
	}

	return false
}

func simulate(last int, motions []string) int {
	knots := make([]knot, last)
	visited := make(map[string]bool)
	visited[fmt.Sprintf("%d;%d", knots[last-1].x, knots[last-1].y)] = true

	for _, motion := range motions {
		var direction rune
		var count int
		_, err := fmt.Sscanf(motion, "%c %d", &direction, &count)

		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < count; i++ {
			switch direction {
			case 'R':
				knots[0].x++
			case 'L':
				knots[0].x--
			case 'U':
				knots[0].y++
			case 'D':
				knots[0].y--

			}

			for i := 0; i < last-1; i++ {
				head := &knots[i]
				tail := &knots[i+1]

				if !isAdjacent(head, tail) {
					if tail.y == head.y {
						if tail.x < head.x {
							tail.x++
						} else {
							tail.x--
						}
					} else if tail.x == head.x {
						if tail.y < head.y {
							tail.y++
						} else {
							tail.y--
						}
					} else {
						if tail.y < head.y {
							tail.y++
						} else {
							tail.y--
						}
						if tail.x < head.x {
							tail.x++
						} else {
							tail.x--
						}
					}
				}
			}

			visited[fmt.Sprintf("%d;%d", knots[last-1].x, knots[last-1].y)] = true
		}
	}

	return len(visited)
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
	motions := make([]string, 0)

	for scanner.Scan() {
		motions = append(motions, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	fmt.Println(simulate(2, motions))
	fmt.Println(simulate(10, motions))
}
