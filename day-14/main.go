package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type coord struct {
	x, y int
}

type grid struct {
	grid [][]byte
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func draw(grid *grid, shape []coord) {
	for i := 0; i < len(shape)-1; i++ {
		if shape[i].x == shape[i+1].x {
			x := shape[i].x
			start := min(shape[i].y, shape[i+1].y)
			end := max(shape[i].y, shape[i+1].y)

			for i := start; i <= end; i++ {
				grid.grid[i][x] = '#'
			}
		} else {
			y := shape[i].y
			start := min(shape[i].x, shape[i+1].x)
			end := max(shape[i].x, shape[i+1].x)

			for i := start; i <= end; i++ {
				grid.grid[y][i] = '#'
			}
		}
	}
}

func print(grid *grid) {
	for i, _ := range grid.grid {
		for j, _ := range grid.grid[i] {
			if grid.grid[i][j] == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%c", grid.grid[i][j])
			}
		}

		fmt.Printf("\n")
	}
}

func sand(grid *grid) int {
	steps := 0

	for {
		coord := coord{x: 500, y: 0}

		if grid.grid[coord.y][coord.x] != 0 {
			break
		}

		for coord.y+1 < len(grid.grid) {
			if grid.grid[coord.y+1][coord.x] == 0 {
				coord.y++
			} else if coord.x-1 >= 0 && grid.grid[coord.y+1][coord.x-1] == 0 {
				coord.y++
				coord.x--
			} else if coord.x+1 < len(grid.grid[0]) && grid.grid[coord.y+1][coord.x+1] == 0 {
				coord.y++
				coord.x++
			} else {
				grid.grid[coord.y][coord.x] = 'o'
				steps++
				break
			}
		}

		if coord.y+1 >= len(grid.grid) {
			break
		}
	}

	return steps
}

func createGrid(shapes [][]coord, xmin, xmax, ymax int) *grid {
	grid := new(grid)
	grid.grid = make([][]byte, 0)

	for i := 0; i <= ymax; i++ {
		grid.grid = append(grid.grid, make([]byte, xmax+200))
	}

	for _, shape := range shapes {
		draw(grid, shape)
	}

	return grid
}

func floor(grid *grid) *grid {
	grid.grid = append(grid.grid, make([]byte, len(grid.grid[0])))
	grid.grid = append(grid.grid, make([]byte, len(grid.grid[0])))
	last := len(grid.grid) - 1

	for i, _ := range grid.grid[len(grid.grid)-1] {
		grid.grid[last][i] = '#'
	}

	return grid
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

	scanner := bufio.NewScanner(input)
	shapes := make([][]coord, 0)
	xmin, xmax, ymax := math.MaxInt, 0, 0

	for scanner.Scan() {
		text := scanner.Text()
		shape := make([]coord, 0)

		for _, coordinate := range strings.Split(text, " -> ") {
			var coord coord
			_, err := fmt.Sscanf(coordinate, "%d,%d", &coord.x, &coord.y)

			if err != nil {
				log.Fatal(err)
			}

			if coord.x > xmax {
				xmax = coord.x
			}
			if coord.y > ymax {
				ymax = coord.y
			}
			if coord.x < xmin {
				xmin = coord.x
			}

			shape = append(shape, coord)
		}

		shapes = append(shapes, shape)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	fmt.Println(xmin, xmax, ymax)
	grid := createGrid(shapes, xmin, xmax, ymax)
	steps := sand(grid)
	fmt.Println(steps)
	print(grid)
	grid = floor(createGrid(shapes, xmin, xmax, ymax))
	steps = sand(grid)
	print(grid)
	fmt.Println(steps)
}
