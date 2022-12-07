package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type node struct {
	name     string
	size     int
	parent   *node
	children []*node
}

var (
	filesystemSize = 70000000
	updateNeededSpace = 30000000
)

func nodeSize(node *node) int {
	size := 0

	if node.children == nil {
		return node.size
	}
	for _, child := range node.children {
		size += nodeSize(child)
	}

	return size
}

func setNodeSize(node *node) {
	node.size = nodeSize(node)

	for _, child := range node.children {
		setNodeSize(child)
	}
}

func sumAtMost(node *node) int {
	total := 0

	if node.children == nil {
		return 0
	}
	if node.size <= 100000 {
		total += node.size
	}

	for _, child := range node.children {
		total += sumAtMost(child)
	}

	return total
}

func freeCandidates(node *node, free int) []int {
	var candidates []int

	if node.children == nil {
		return candidates
	}

	if free + node.size >= updateNeededSpace {
		candidates = append(candidates, node.size)
	}

	for _, child := range node.children {
		if c := freeCandidates(child, free); c != nil {
			candidates = append(candidates, c...)
		}
	}

	return candidates
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
	root := &node{name: "/", size: 0, parent: nil, children: nil}
	current := root
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		text := scanner.Text()

		var directory string
		_, err := fmt.Sscanf(text, "$ cd %s", &directory)

		if err == nil {
			switch directory {
			case "/":
				continue
			case "..":
				current = current.parent
			default:
				child := &node{name: directory, size: 0, parent: current, children: nil}
				current.children = append(current.children, child)
				current = child
			}

			continue
		}

		var file string
		var size int
		_, err = fmt.Sscanf(text, "%d %s", &size, &file)

		if err == nil {
			child := &node{name: file, size: size, parent: current, children: nil}
			current.children = append(current.children, child)
		}
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	setNodeSize(root)
	fmt.Printf("Sum of at most 100000: %d\n", sumAtMost(root))
	candidates := freeCandidates(root, filesystemSize - root.size)
	sort.Ints(candidates)
	fmt.Printf("Smallest candidates to free: %d\n", candidates[0])
}
