package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(stacks []Stack, instructions []string, part int) string {
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
			// create temporary stack to place the items on, to that they retain their
			// original order when placing them on the destination stack
			tempStack := Stack{}
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
		topOfEachStack = append(topOfEachStack, stack.elements[len(stack.elements)-1])
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
	stacks := make([]Stack, numStacks)

	for _, line := range cratePlan {
		for i := 0; 4*i+1 < len(line); i++ {
			index := 4*i + 1
			if string(line[index]) == " " {
				continue
			}
			stacks[i].Push(string(line[index]))
		}

	}
	// since items were added from top to bottom, we need to reverse
	for _, stack := range stacks {
		stack.Reverse()
	}

	// In order for solve() to return the proper result, I need to comment-out either part 1 or part 2.
	// The reason for this lies in the nature of slices. Even though the slice is passed by value
	// it still contains a pointer to the underlying array data structure.
	// Modifications on that array are hence reflected in the slice as well.
	// That means that the stacks variable is always modified when solve() is called.

	// fmt.Println(solve(stacks, instructions, 1))
	fmt.Println(solve(stacks, instructions, 2))
}
