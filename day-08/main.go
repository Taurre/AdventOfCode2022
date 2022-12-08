package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func isVisible(grid [][]int, y, x int) bool {
	edge := len(grid) - 1
	visible := 4

	if y == 0 || x == 0 || y == edge || x == edge {
		return true
	}

	for top := y - 1; top >= 0; top-- {
		if grid[top][x] >= grid[y][x] {
			visible--
			break
		}
	}

	for bottom := y + 1; bottom <= edge; bottom++ {
		if grid[bottom][x] >= grid[y][x] {
			visible--
			break
		}
	}

	for left := x - 1; left >= 0; left-- {
		if grid[y][left] >= grid[y][x] {
			visible--
			break
		}
	}

	for right := x + 1; right <= edge; right++ {
		if grid[y][right] >= grid[y][x] {
			visible--
			break
		}
	}

	if visible == 0 {
		return false
	} else {
		return true
	}
}

func scenicScore(grid [][]int, y, x int) int {
	edge := len(grid) - 1
	score := make([]int, 4)

	for top := y - 1; top >= 0; top-- {
		score[0]++

		if grid[top][x] >= grid[y][x] {
			break
		}
	}

	for bottom := y + 1; bottom <= edge; bottom++ {
		score[1]++

		if grid[bottom][x] >= grid[y][x] {
			break
		}
	}

	for left := x - 1; left >= 0; left-- {
		score[2]++

		if grid[y][left] >= grid[y][x] {
			break
		}
	}

	for right := x + 1; right <= edge; right++ {
		score[3]++

		if grid[y][right] >= grid[y][x] {
			break
		}
	}

	return score[0] * score[1] * score[2] * score[3]
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
	grid := make([][]int, 0)

	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()
		line := make([]int, 0)

		for _, c := range text {
			line = append(line, int(c-'0'))
		}

		grid = append(grid, line)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	count := 0
	highestScore := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if isVisible(grid, i, j) {
				count++
			}

			score := scenicScore(grid, i, j)

			if score > highestScore {
				highestScore = score
			}
		}
	}

	fmt.Println(count)
	fmt.Println(highestScore)
}
