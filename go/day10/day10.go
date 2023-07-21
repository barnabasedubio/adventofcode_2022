package main

import (
	"aoc2022/shared"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getToAdd(operationMap map[int]string, key int) int {
	if operationMap[key] == "noop" {
		return 0
	}
	valStr := strings.Split(operationMap[key], " ")[1]
	val, _ := strconv.Atoi(valStr)
	return val
}

func part1(inputList []string) ([]int, int) {
	instructionStack := shared.Stack[string]{}
	instructionStack.Elements = inputList
	instructionStack.Reverse()
	xHistory := []int{1}
	for len(instructionStack.Elements) > 0 {
		newInstruction, _ := instructionStack.Pop()
		xHistory = append(xHistory, xHistory[len(xHistory)-1])
		if strings.HasPrefix(newInstruction, "addx ") {
			toAdd, _ := strconv.Atoi(strings.Split(newInstruction, " ")[1])
			newX := xHistory[len(xHistory)-1] + toAdd
			xHistory = append(xHistory, newX)
		}
	}
	sum := 0
	cycles := []int{20, 60, 100, 140, 180, 220}
	for _, v := range cycles {
		sum += v * xHistory[v-1]
	}
	return xHistory, sum
}

func part2(xHistory []int) string {
	stringBuilder := []string{}
	for i := 0; i < 240; i++ {
		left := xHistory[i]-1
		center := xHistory[i]
		right := xHistory[i]+1
		if i%40 == left || i%40 == center || i%40 == right {
			stringBuilder = append(stringBuilder, "#")
		} else {
			stringBuilder = append(stringBuilder, ".")
		}
		if i%40 == 39 {
			stringBuilder = append(stringBuilder, "\n")
		}
	}
	return strings.Join(stringBuilder, "")
}

func main() {
	input, _ := os.ReadFile("day10.in")
	inputList := strings.Split(string(input), "\n")
	xHistory, xSum := part1(inputList)
	fmt.Println(xSum)
	fmt.Println(part2(xHistory))
}
