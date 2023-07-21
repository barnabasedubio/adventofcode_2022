package main

import (
	"aoc2022/shared"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type directory struct {
	id   string
	name string
}

func solve(inputList []string, part int) int {
	directoryLineage := shared.Stack[directory]{}
	directoryMap := map[directory]int{}

	for _, line := range inputList {
		switch {
		case strings.HasPrefix(line, "dir ") || line == "$ ls":
			continue
		case line == "$ cd ..":
			directoryLineage.Pop()
		case strings.HasPrefix(line, "$ cd "):
			target := strings.Split(line, " ")[2]
			directoryLineage.Push(directory{
				id:   shared.GenerateUID(8),
				name: target,
			})
		default:
			size, _ := strconv.Atoi((strings.Split(line, " ")[0]))
			for _, directory := range directoryLineage.Elements {
				directoryMap[directory] += size
			}
		}
	}
	switch part {
	case 1:
		total := 0
		for _, val := range directoryMap {
			if val <= 100000 {
				total += val
			}
		}
		return total
	case 2:
		totalSpace := 70000000
		requiredSpace := 30000000
		availableSpace := -1
		for k, v := range directoryMap {
			if k.name == "/" {
				availableSpace = totalSpace - v
				break
			}
		}
		if availableSpace < 0 {
			panic("Availabe space was not set properly")
		}
		toDeleteSpace := requiredSpace - availableSpace
		candidates := []int{}
		for _, v := range directoryMap {
			if v >= toDeleteSpace {
				candidates = append(candidates, v)
			}
		}
		sort.Ints(candidates)
		return candidates[0]
	}
	return -1
}

func main() {
	input, _ := os.ReadFile("day7.in")
	inputList := strings.Split(string(input), "\n")
	fmt.Println(solve(inputList, 1))
	fmt.Println(solve(inputList, 2))
} 