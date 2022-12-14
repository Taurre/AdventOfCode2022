package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"unicode"
)

const (
	less = iota
	equal
	greater
)

type cmp int
type packets []string

func (packets packets) Len() int {
	return len(packets)
}

func (packets packets) Less(i, j int) bool {
	cmp := compare(packets[i], packets[j])

	if cmp == greater {
		return false
	}

	return true
}

func (packets packets) Swap(i, j int) {
	packets[i], packets[j] = packets[j], packets[i]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func split(list string) []string {
	level := 0
	start := 0
	s := make([]string, 0)

	for i := 0; i < len(list); i++ {
		switch list[i] {
		case '[':
			level++
		case ']':
			level--
		case ',':
			if level == 0 {
				s = append(s, list[start:i])
				start = i+1
			}
		}
	}

	if len(list[start:]) > 0 {
		s = append(s, list[start:])
	}

	return s
}

func compare(a, b string) cmp {
	a = unwrap(a)
	b = unwrap(b)

	if a == "" || b == "" {
		if a == "" && b == "" {
			return equal
		} else if a == "" {
			return less
		} else {
			return greater
		}
	}

	s := [][]string{split(a), split(b)}
	length := min(len(s[0]), len(s[1]))

	for i := 0; i < length; i++ {
		left, right := s[0][i], s[1][i]
		var cmp cmp = equal

		if left[0] == '[' && right[0] == '[' {
			cmp = compare(left, right)
		} else if unicode.IsDigit(rune(left[0])) && unicode.IsDigit(rune(right[0])) {
			a, err := strconv.Atoi(left)

			if err != nil {
				log.Fatal(err)
			}

			b, err := strconv.Atoi(right)

			if err != nil {
				log.Fatal(err)
			}

			if a < b {
				return less
			} else if a > b {
				return greater
			}
		} else if unicode.IsDigit(rune(left[0])) {
			cmp = compare("["+left+"]", right)
		} else {
			cmp = compare(left, "["+right+"]")
		}

		if cmp == less || cmp == greater {
			return cmp
		}
	}

	if len(s[0]) < len(s[1]) {
		return less
	} else if len(s[0]) > len(s[1]) {
		return greater
	}

	return equal
}

func unwrap(s string) string {
	level := 1
	i := 1

	for ; i < len(s) && level > 0; i++ {
		switch s[i] {
		case '[':
			level++
		case ']':
			level--
		}

		if level == 0 {
			break
		}
	}

	return s[1:i]
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
	packets := make(packets, 0)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			continue
		}

		packets = append(packets, text)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	indice := 1
	sum := 0

	for i := 0; i < len(packets); i += 2 {
		if (compare(packets[i], packets[i+1]) == less) {
			sum += indice
		}

		indice++
	}

	fmt.Println(sum)
	packets = append(packets, "[[2]]", "[[6]]")
	sort.Stable(packets)
	key := 1

	for i := 0; i < len(packets); i++ {
		if packets[i] == "[[2]]" || packets[i] == "[[6]]" {
			key *= i+1
		}
	}

	fmt.Println(key)
}
