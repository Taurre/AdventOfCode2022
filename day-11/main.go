package main

import (
	"fmt"
	"sort"
)

type operation struct {
	operator rune
	value    int
}

type monkey struct {
	items     []int
	operation operation
	divisible int
	ifTrue    int
	ifFalse   int
	inspect   int
}

var (
	monkeys = []monkey {
		monkey{
			items: []int{66, 59, 64, 51},
			operation: operation{operator: '*', value: 3},
			divisible: 2,
			ifTrue: 1,
			ifFalse: 4,
			inspect: 0,
		},
		monkey{
			items: []int{67, 61},
			operation: operation{operator: '*', value: 19},
			divisible: 7,
			ifTrue: 3,
			ifFalse: 5,
			inspect: 0,
		},
		monkey{
			items: []int{86, 93, 80, 70, 71, 81, 56},
			operation: operation{operator: '+', value: 2},
			divisible: 11,
			ifTrue: 4,
			ifFalse: 0,
			inspect: 0,
		},
		monkey{
			items: []int{94},
			operation: operation{operator: '*', value: 0},
			divisible: 19,
			ifTrue: 7,
			ifFalse: 6,
			inspect: 0,
		},
		monkey{
			items: []int{71, 92, 64},
			operation: operation{operator: '+', value: 8},
			divisible: 3,
			ifTrue: 5,
			ifFalse: 1,
			inspect: 0,
		},
		monkey{
			items: []int{58, 81, 92, 75, 56},
			operation: operation{operator: '+', value: 6},
			divisible: 5,
			ifTrue: 3,
			ifFalse: 6,
			inspect: 0,
		},
		monkey{
			items: []int{82, 98, 77, 94, 86, 81},
			operation: operation{operator: '+', value: 7},
			divisible: 17,
			ifTrue: 7,
			ifFalse: 2,
			inspect: 0,
		},
		monkey{
			items: []int{54, 95, 70, 93, 88, 93, 63, 50},
			operation: operation{operator: '+', value: 4},
			divisible: 13,
			ifTrue: 2,
			ifFalse: 0,
			inspect: 0,
		},
	}
)

func round(monkeys []monkey) {
	for i, _ := range monkeys {
		for _, item := range monkeys[i].items {
			operator := monkeys[i].operation.operator
			value := monkeys[i].operation.value

			switch operator {
			case '+':
				if value == 0 {
					value = item + item
				} else {
					value += item
				}
			case '*':
				if value == 0 {
					value = item * item
				} else {
					value *= item
				}
			}

			value /= 3

			if value % monkeys[i].divisible == 0 {
				target := monkeys[i].ifTrue
				monkeys[target].items = append(monkeys[target].items, value)
			} else {
				target := monkeys[i].ifFalse
				monkeys[target].items = append(monkeys[target].items, value)
			}

			monkeys[i].inspect++
		}

		monkeys[i].items = monkeys[i].items[:0]	
	}
}

func main() {
	for i := 0; i < 20; i++ {
		round(monkeys)
	}

	inspects := make([]int, 0, len(monkeys))

	for _, monkey := range monkeys {
		inspects = append(inspects, monkey.inspect)
	}

	sort.Ints(inspects)
	last := len(inspects)-1
	fmt.Println(inspects[last-1] * inspects[last])
}
