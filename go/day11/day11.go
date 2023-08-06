package main

import (
	"aoc2022/shared"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInformation(monkeyInfo string) (shared.Queue[int], string, int, int, int, int) {
	startingItems := shared.Queue[int]{}
	operation := "" // add, mult, sub, div
	operand := 0
	divisibleBy := 0
	ifTrue := 0
	ifFalse := 0
	monkeyInfoList := strings.Split(monkeyInfo, "\n")
	// parse starting items by populating queue
	queueElementsStr := strings.Split(monkeyInfoList[1], ": ")[1]
	queueElements := strings.Split(queueElementsStr, ", ")
	for _, val := range queueElements {
		elemAsInt, _ := strconv.Atoi(val)
		startingItems.Push(elemAsInt)
	}
	// parse operation
	operationStr := strings.Split(monkeyInfoList[2], " = ")[1]
	symbol := strings.Split(operationStr, " ")[1]
	switch symbol {
	case "*":
		operation = "mult"
	case "+":
		operation = "add"
	case "-":
		operation = "sub"
	case "/":
		operation = "div"
	default:
		errorStr := fmt.Sprintf("unknown operand: %v", symbol)
		panic(errorStr)
	}
	// parse operand
	operand, _ = strconv.Atoi(strings.Split(operationStr, " ")[2])
	// parse divisble by
	divisibleBy, _ = strconv.Atoi(strings.Split(monkeyInfoList[3], " ")[5])
	// parse test true case
	ifTrue, _ = strconv.Atoi(strings.Split(monkeyInfoList[4], " ")[9])
	// parse test false case
	ifFalse, _ = strconv.Atoi(strings.Split(monkeyInfoList[5], " ")[9])
	return startingItems, operation, operand, divisibleBy, ifTrue, ifFalse
}

func part1(monkeyInstructions []string) {
	// round := 0
	monkeyMap := map[int]map[string]any{}
	// initialize monkeyData
	for index, instructions := range monkeyInstructions {
		startingItems, operation, operand, divisibleBy, ifTrue, ifFalse := parseInformation(instructions)
		innerMonkeyMap := map[string]any{}
		innerMonkeyMap["queue"] = startingItems
		innerMonkeyMap["operation"] = operation
		innerMonkeyMap["operand"] = operand
		innerMonkeyMap["divisibleBy"] = divisibleBy
		innerMonkeyMap["ifTrue"] = ifTrue
		innerMonkeyMap["ifFalse"] = ifFalse
		innerMonkeyMap["inspected"] = 0
		monkeyMap[index] = innerMonkeyMap
	}
	fmt.Println(monkeyMap)
	// for round <= 20 {
	// round++
	// }
}

func main() {
	input, _ := os.ReadFile("day11.test.in")
	monkeyInstructions := strings.Split(string(input), "\n\n")
	part1(monkeyInstructions)
}
