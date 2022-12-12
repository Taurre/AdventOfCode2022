package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

type square struct {
	letter    rune
	elevation int
	distance  int
	visited   bool
}

type coord struct {
	x, y int
}

type heightmap [][]square

func (heightmap heightmap) isValid(coord coord) bool {
	if coord.x < 0 || coord.y < 0 {
		return false
	}
	if coord.x >= len(heightmap) || coord.y >= len(heightmap[0]) {
		return false
	}

	return true
}

func shortest(hm heightmap, from, to coord) int {
	heightmap := make(heightmap, len(hm))

	for i := 0; i < len(hm); i++ {
		heightmap[i] = make([]square, len(hm[i]))
		copy(heightmap[i], hm[i])
	}

	heightmap[from.x][from.y].distance = 0
	heightmap[from.x][from.y].visited = true
	nodes := append([]coord(nil), from)

	for !heightmap[to.x][to.y].visited && len(nodes) > 0 {
		next := make([]coord, 0)

		for _, node := range nodes {
			for _, adjacent := range []coord{coord{-1, 0}, coord{1, 0}, coord{0, -1}, coord{0, 1}} {
				adjacent.x += node.x
				adjacent.y += node.y

				if !heightmap.isValid(adjacent) {
					continue
				}
				if heightmap[adjacent.x][adjacent.y].visited {
					continue
				}
				if heightmap[node.x][node.y].elevation+1 < heightmap[adjacent.x][adjacent.y].elevation {
					continue
				}
				if heightmap[adjacent.x][adjacent.y].distance > heightmap[node.x][node.y].distance+1 {
					heightmap[adjacent.x][adjacent.y].distance = heightmap[node.x][node.y].distance + 1
				}

				next = append(next, adjacent)
				heightmap[adjacent.x][adjacent.y].visited = true
			}
		}

		nodes = next
	}

	return heightmap[to.x][to.y].distance
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

	heightmap := make(heightmap, 0)
	scanner := bufio.NewScanner(input)
	start := coord{x: 0, y: 0}
	end := coord{x: 0, y: 0}

	for i := 0; scanner.Scan(); i++ {
		heightmap = append(heightmap, make([]square, 0))
		text := scanner.Text()

		for _, letter := range text {
			elevation := int(letter - 'a')

			switch letter {
			case 'S':
				elevation = 'a' - 'a'
				start.x = i
				start.y = len(heightmap[i])
			case 'E':
				elevation = 'z' - 'a'
				end.x = i
				end.y = len(heightmap[i])
			}

			square := square{letter: letter, elevation: elevation, distance: math.MaxInt32, visited: false}
			heightmap[i] = append(heightmap[i], square)
		}
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	fmt.Println(shortest(heightmap, start, end))
	distances := make([]int, 0)

	for i, _ := range heightmap {
		for j, _ := range heightmap[i] {
			if heightmap[i][j].letter == 'a' {
				start.x = i
				start.y = j
				distances = append(distances, shortest(heightmap, start, end))
			}
		}
	}

	sort.Ints(distances)
	fmt.Println(distances[0])
}
