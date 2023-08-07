package main

import (
	"aoc2022/shared"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInformation(monkeyInfo string) (shared.Queue[int64], string, int64, int64, int, int) {
	startingItems := shared.Queue[int64]{}
	operation := "" // add, mult, sub, div
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
	case "-":
		operation = "sub"
	case "/":
		operation = "div"
	default:
		errorStr := fmt.Sprintf("unknown operand: %v", symbol)
		panic(errorStr)
	}
	// parse operand
	if strings.Split(operationStr, " ")[2] == "old" {
		// placeholder that we should use the old value as operand. This will be handled in the
		// getWorryLevel() function.
		operand = 999
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

func getWorryLevel(value int64, operation string, operand int64) int64 {
	val := int64(0)
	if operand == 999 {
		operand = value
	}
	switch operation {
	case "mult":
		val = int64(value * operand)
	case "add":
		val = int64(value + operand)
	case "div":
		val = int64(value / operand)
	case "sub":
		val = int64(value - operand)
	default:
		panic("Unknown operation")
	}
	if val == 0 {
		errorString := fmt.Sprintf("value was 0.\nvalue: %v, operation: %v, operand: %v", value, operation, operand)
		panic(errorString)
	}
	return val
}

func solve(monkeyInstructions []string, part int) int {
	monkeyMap := map[int]map[string]any{}
	// initialize monkeyData
	indexList := []int{}
	for index, instructions := range monkeyInstructions {
		indexList = append(indexList, index)
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
	round := 0
	for round < 20 {
		for _, i := range indexList {
			queue := monkeyMap[i]["queue"].(shared.Queue[int64])
			operation := monkeyMap[i]["operation"].(string)
			operand := monkeyMap[i]["operand"].(int64)
			divisibleBy := monkeyMap[i]["divisibleBy"].(int64)
			ifTrue := monkeyMap[i]["ifTrue"].(int)
			ifFalse := monkeyMap[i]["ifFalse"].(int)

			if queue.IsEmpty() {
				continue
			}
			for !queue.IsEmpty() {
				// monkey inspects item
				currentItem, _ := queue.Pop()
				monkeyMap[i]["queue"] = queue
				numTimesInspected := monkeyMap[i]["inspected"].(int)
				numTimesInspected++
				monkeyMap[i]["inspected"] = numTimesInspected
				worryLevel := getWorryLevel(currentItem, operation, operand)

				worryLevel /= 3

				if worryLevel%divisibleBy == 0 {
					if targetQueue, ok := monkeyMap[ifTrue]["queue"].(shared.Queue[int64]); ok {
						targetQueue.Push(worryLevel)
						monkeyMap[ifTrue]["queue"] = targetQueue
					} else {
						panic("Target queue is not of type Queue")
					}
				} else {
					if targetQueue, ok := monkeyMap[ifFalse]["queue"].(shared.Queue[int64]); ok {
						targetQueue.Push(worryLevel)
						monkeyMap[ifFalse]["queue"] = targetQueue
					} else {
						panic("Target queue is not of type Queue")
					}
				}
			}
		}
		round++
	}
	// monkeyActivity is a list containing # times monkey <index> inspected an item
	monkeyActivity := []int{}
	for _, id := range indexList {
		fmt.Printf("Monkey %v\n", id)
		fmt.Println(monkeyMap[id]["queue"])
		fmt.Println(monkeyMap[id]["inspected"])
		inspected := monkeyMap[id]["inspected"].(int)
		monkeyActivity = append(monkeyActivity, inspected)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(monkeyActivity)))
	fmt.Println(monkeyActivity)
	return monkeyActivity[0] * monkeyActivity[1]
}

func main() {
	input, _ := os.ReadFile("day11.in")
	monkeyInstructions := strings.Split(string(input), "\n\n")
	fmt.Println(solve(monkeyInstructions, 1))
}
