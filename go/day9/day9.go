package main

import (
	"aoc2022/shared"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateNextHeadPosition(headPath []position, direction string) position {
	last := headPath[len(headPath)-1]
	next := position{}
	switch direction {
	case "L":
		next = position{last.x - 1, last.y}
	case "R":
		next = position{last.x + 1, last.y}
	case "U":
		next = position{last.x, last.y + 1}
	case "D":
		next = position{last.x, last.y - 1}
	}
	return next
}

func calculateHeadPath(inputList []string) []position {
	headPath := []position{{0, 0}}
	for _, instruction := range inputList {
		direction := strings.Split(instruction, " ")[0]
		amount, _ := strconv.Atoi(strings.Split(instruction, " ")[1])
		for i := 0; i < amount; i++ {
			nextHeadPosition := calculateNextHeadPosition(headPath, direction)
			headPath = append(headPath, nextHeadPosition)
		}
	}
	return headPath
}

func calculateTailPath(headPath []position) []position {
	tailPath := []position{{0, 0}}
	for i, nextHead := range headPath {
		currTail := tailPath[len(tailPath)-1]
		switch {
		case i == 0:
			continue

		case nextHead == currTail:
			// overlap, no change required
			tailPath = append(tailPath, currTail)

		case (abs(nextHead.x-currTail.x) == 1 && (nextHead.y-currTail.y == 0)) ||
			(abs(nextHead.y-currTail.y) == 1 && (nextHead.x-currTail.x == 0)):
			// direct neighbor, no change required
			tailPath = append(tailPath, currTail)

		case abs(nextHead.x-currTail.x)*abs(nextHead.y-currTail.y) == 1:
			// diagonal neighbor, no change required
			tailPath = append(tailPath, currTail)

		case abs(nextHead.x-currTail.x) == 2 && nextHead.y-currTail.y == 0:
			// move along x axis
			newX := int((nextHead.x + currTail.x) / 2)
			tailPath = append(tailPath, position{newX, currTail.y})

		case abs(nextHead.y-currTail.y) == 2 && nextHead.x-currTail.x == 0:
			// move along y axis
			newY := int((nextHead.y + currTail.y) / 2)
			tailPath = append(tailPath, position{currTail.x, newY})

		default:
			deltaX := (nextHead.x - currTail.x)
			deltaY := (nextHead.y - currTail.y)
			var newX int
			var newY int
			if deltaX < 0 {
				newX = currTail.x - 1
			} else if deltaX > 0 {
				newX = currTail.x + 1
			}
			if deltaY < 0 {
				newY = currTail.y - 1
			} else if deltaY > 0 {
				newY = currTail.y + 1
			}
			tailPath = append(tailPath, position{newX, newY})
		}
	}
	return tailPath
}

func solve(inputList []string, part int) int {
	headPath := calculateHeadPath(inputList)
	tailPath := calculateTailPath(headPath)
	switch part {
	case 1:
		return len(shared.Set(tailPath))
	case 2:
		for i := 0; i < 8; i++ {
			tailPath = calculateTailPath(tailPath)
		}
	}
	return len(shared.Set(tailPath))

}

func main() {
	input, _ := os.ReadFile("day9.in")
	inputList := strings.Split(string(input), "\n")
	fmt.Println(solve(inputList, 1))
	fmt.Println(solve(inputList, 2))
}
