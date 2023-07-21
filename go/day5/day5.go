package main

import (
	"aoc2022/shared"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(stacks []shared.Stack[string], instructions []string, part int) string {
	topOfEachStack := []string{}
	for _, instruction := range instructions {
		instructionList := strings.Split(instruction, " ")
		source, _ := strconv.Atoi(instructionList[3])
		dest, _ := strconv.Atoi(instructionList[5])
		amount, _ := strconv.Atoi(instructionList[1])

		switch part {
		case 1:
			for i := 0; i < amount; i++ {
				popped, _ := stacks[source-1].Pop()
				stacks[dest-1].Push(popped)
			}
		case 2:
			// create temporary stack to place the items on in order to maintain
			// original order when placing them on the destination stack
			tempStack := shared.Stack[string]{}
			for i := 0; i < amount; i++ {
				popped, _ := stacks[source-1].Pop()
				tempStack.Push(popped)
			}
			for i := 0; i < amount; i++ {
				popped, _ := tempStack.Pop()
				stacks[dest-1].Push(popped)
			}
		}
	}
	for _, stack := range stacks {

		topOfEachStack = append(topOfEachStack, stack.Elements[len(stack.Elements)-1])
	}
	return strings.Join(topOfEachStack, "")
}

func main() {
	input, _ := os.ReadFile("day5.in")
	cratePlan := strings.Split((strings.Split(string(input), "\n\n"))[0], "\n")
	stackLabels := strings.Trim(cratePlan[len(cratePlan)-1], " ")
	numStacks, _ := strconv.Atoi((string(stackLabels[len(stackLabels)-1])))
	cratePlan = cratePlan[:len(cratePlan)-1]

	instructions := strings.Split((strings.Split(string(input), "\n\n"))[1], "\n")

	// I will need to make identical versions of the stackList, for part 1 and part 2.
	// calling solve() using a single slice for both parts is not going to work since
	// solve() modifies the underlying array the slice is pointing to
	stacks1 := make([]shared.Stack[string], numStacks)
	stacks2 := make([]shared.Stack[string], numStacks)

	// extract the arrangement of the crates from the input text
	for _, line := range cratePlan {
		for i := 0; 4*i+1 < len(line); i++ {
			index := 4*i + 1
			if string(line[index]) == " " {
				continue
			}
			stacks1[i].Push(string(line[index]))
			stacks2[i].Push(string(line[index]))
		}

	}
	// since items were added from top to bottom, we need to reverse
	for i := 0; i < len(stacks1); i++ {
		stacks1[i].Reverse()
		stacks2[i].Reverse()
	}

	fmt.Println(solve(stacks1, instructions, 1))
	fmt.Println(solve(stacks2, instructions, 2))
}
