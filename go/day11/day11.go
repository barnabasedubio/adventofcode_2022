package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func part1(monkeyInstructions []string) int64 {
	monkeys := []Monkey{}
	for i, instructions := range monkeyInstructions {
		m := Monkey{}
		m.Id = i
		startingItems, operation, operand, divisibleBy, trueCase, falseCase := ParseInformation(instructions)
		m.Items = startingItems
		m.Operation = operation
		m.Operand = operand
		m.DivisibleBy = divisibleBy
		m.TrueCase = trueCase
		m.FalseCase = falseCase
		monkeys = append(monkeys, m)
	}
	round := 0
	for round < 20 {
		for i := range monkeys {
			if !monkeys[i].HasItems() {
				continue
			}
			for monkeys[i].HasItems() {
				monkeys[i].Inspect()
				monkeys[i].ReduceWorryLevel()
				targetMonkeyId := monkeys[i].Test()
				monkeys[i].Throw(&monkeys[targetMonkeyId])
			}
		}
		round++
	}
	for _, monkey := range monkeys {
		fmt.Println(monkey)
	}
	monkeyActivity := []int{}
	for _, monkey := range monkeys {
		monkeyActivity = append(monkeyActivity, int(monkey.TimesInspected))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(monkeyActivity)))
	monkeyBusiness := monkeyActivity[0] * monkeyActivity[1]
	return int64(monkeyBusiness)
}

func main() {
	input, _ := os.ReadFile("day11.in")
	monkeyInstructions := strings.Split(string(input), "\n\n")
	fmt.Println(part1(monkeyInstructions))
}
