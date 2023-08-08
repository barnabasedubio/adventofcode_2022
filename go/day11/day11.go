package main

import (
	"aoc2022/shared"
	"fmt"
	"os"
	"sort"
	"strings"
)

func calculateMonkeyBusiness(monkeys []Monkey) int64 {
	monkeyActivity := []int{}
	for _, monkey := range monkeys {
		monkeyActivity = append(monkeyActivity, int(monkey.TimesInspected))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(monkeyActivity)))
	return int64(monkeyActivity[0] * monkeyActivity[1])
}

func calculateModMonkeyBusiness(monkeys []ModMonkey) int64 {
	monkeyActivity := []int{}
	for _, monkey := range monkeys {
		monkeyActivity = append(monkeyActivity, int(monkey.TimesInspected))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(monkeyActivity)))
	return int64(monkeyActivity[0] * monkeyActivity[1])
}

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
	return calculateMonkeyBusiness(monkeys)
}

func part2(monkeyInstructions []string) int64 {
	modMonkeys := []ModMonkey{}
	divisors := []int64{}
	startingItemsList := [][]int64{}
	for i, instructions := range monkeyInstructions {
		m := ModMonkey{}
		m.Id = i
		startingItems, operation, operand, divisibleBy, trueCase, falseCase := ParseInformation(instructions)
		startingItemsList = append(startingItemsList, startingItems.Elements)
		m.Operation = operation
		m.Operand = operand
		divisors = append(divisors, divisibleBy)
		m.TrueCase = trueCase
		m.FalseCase = falseCase
		modMonkeys = append(modMonkeys, m)
	}
	for i := range modMonkeys {
		modMonkeys[i].Divisors = divisors
	}
	for i, startingItems := range startingItemsList {
		startingItemsQueue := shared.Queue[[]int64]{}
		for _, startingItem := range startingItems {
			queueList := []int64{}
			for _, divisor := range divisors {
				queueList = append(queueList, startingItem%divisor)
			}
			startingItemsQueue.Push(queueList)
		}
		modMonkeys[i].Items = startingItemsQueue
	}
	round := 0
	for round < 10000 {
		for i := range modMonkeys {
			if !modMonkeys[i].HasItems() {
				continue
			}
			for modMonkeys[i].HasItems() {
				modMonkeys[i].Inspect()
				targetModMonkeyId := modMonkeys[i].Test()
				modMonkeys[i].Throw(&modMonkeys[targetModMonkeyId])
			}
		}
		round++
	}
	return calculateModMonkeyBusiness(modMonkeys)

}

func main() {
	input, _ := os.ReadFile("day11.in")
	monkeyInstructions := strings.Split(string(input), "\n\n")
	fmt.Println(part1(monkeyInstructions))
	fmt.Println(part2(monkeyInstructions))
}
