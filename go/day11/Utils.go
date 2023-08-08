package main

import (
	"aoc2022/shared"
	"fmt"
	"strconv"
	"strings"
)

func ParseInformation(monkeyInfo string) (shared.Queue[int64], string, int64, int64, int, int) {
	startingItems := shared.Queue[int64]{}
	operation := ""
	operand := int64(0)
	divisibleBy := int64(0)
	ifTrue := 0
	ifFalse := 0
	monkeyInfoList := strings.Split(monkeyInfo, "\n")
	// parse starting items by populating queue
	queueElementsStr := strings.Split(monkeyInfoList[1], ": ")[1]
	queueElements := strings.Split(queueElementsStr, ", ")
	for _, val := range queueElements {
		elemAsInt, _ := strconv.Atoi(val)
		startingItems.Push(int64(elemAsInt))
	}
	// parse operation
	operationStr := strings.Split(monkeyInfoList[2], " = ")[1]
	symbol := strings.Split(operationStr, " ")[1]
	switch symbol {
	case "*":
		operation = "mult"
	case "+":
		operation = "add"
	default:
		panic(fmt.Sprintf("Unknown operand: %v", symbol))
	}
	// parse operand
	if strings.Split(operationStr, " ")[2] == "old" {
		// placeholder that we should use the old value as operand. This will be handled in the
		// getWorryLevel() function.
		operand = 999
		operation = "square"
	} else {
		newOperand, _ := strconv.Atoi(strings.Split(operationStr, " ")[2])
		operand = int64(newOperand)
	}
	// parse divisble by
	newDivisbleBy, _ := strconv.Atoi(strings.Split(monkeyInfoList[3], " ")[5])
	divisibleBy = int64(newDivisbleBy)
	// parse test true case
	ifTrue, _ = strconv.Atoi(strings.Split(monkeyInfoList[4], " ")[9])
	// parse test false case
	ifFalse, _ = strconv.Atoi(strings.Split(monkeyInfoList[5], " ")[9])
	return startingItems, operation, operand, divisibleBy, ifTrue, ifFalse
}
