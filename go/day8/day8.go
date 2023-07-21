package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}

func getVerticalTrees(inputMap [][]int, j int) []int {
	lst := []int{}
	for _, val := range inputMap {
		lst = append(lst, val[j])
	}
	return lst
}

func isLargerThanAll(x int, values []int) bool {
	ret := true
	for _, val := range values {
		if val >= x {
			ret = false
		}
	}
	return ret
}

func reverseList(list []int) []int {
	reversed := []int{}
	for _, v := range list {
		reversed = append(reversed, v)
	}
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	return reversed
}

func isVisible(i int, j int, inputMap [][]int) bool {
	west := inputMap[i][:j]
	east := inputMap[i][j+1 : len(inputMap[i])]
	north := getVerticalTrees(inputMap, j)[:i]
	south := getVerticalTrees(inputMap, j)[i+1 : len(inputMap)]

	currVal := inputMap[i][j]

	return isLargerThanAll(currVal, west) ||
		isLargerThanAll(currVal, east) ||
		isLargerThanAll(currVal, north) ||
		isLargerThanAll(currVal, south)
}

func numberOfTreesVisible(value int, direction []int) int {
	// at least one tree is visible to each direction
	// edge trees can not be counted since their scenic values would
	// alway be 0.
	visible := 0
	for _, tree := range direction {
		// fmt.Println(tree)
		if tree < value {
			visible++
		} else if tree >= value {
			visible++
			break
		}
	}
	return max(visible, 1)
}

func calculateScenicScore(i int, j int, inputMap [][]int) int {
	west := inputMap[i][:j]
	east := inputMap[i][j+1 : len(inputMap[i])]
	north := getVerticalTrees(inputMap, j)[:i]
	south := getVerticalTrees(inputMap, j)[i+1 : len(inputMap)]
	// since we are looking "outward" at trees from the position of a given tree
	// trees to the north and west need to be reversed
	west = reverseList(west)
	north = reverseList(north)

	currVal := inputMap[i][j]
	westVisible := numberOfTreesVisible(currVal, west)
	eastVisible := numberOfTreesVisible(currVal, east)
	northVisible := numberOfTreesVisible(currVal, north)
	southVisible := numberOfTreesVisible(currVal, south)

	return westVisible * eastVisible * northVisible * southVisible
}

func part1(inputMap [][]int) int {
	// trees on the edge are visible by default
	visible := len(inputMap[0])*2 + (len(inputMap)-2)*2
	for i := 1; i < len(inputMap)-1; i++ {
		for j := 1; j < len(inputMap[0])-1; j++ {
			if isVisible(i, j, inputMap) {
				visible++
			}
		}
	}
	return visible
}

func part2(inputMap [][]int) int {
	scenic_score := 0
	for i := 1; i < len(inputMap)-1; i++ {
		for j := 1; j < len(inputMap[0])-1; j++ {
			currentScenicScore := calculateScenicScore(i, j, inputMap)
			scenic_score = max(scenic_score, currentScenicScore)
		}
	}
	return int(scenic_score)
}

func main() {
	input, _ := os.ReadFile("day8.in")
	inputList := strings.Split(string(input), "\n")
	inputMap := [][]int{}
	for _, row := range inputList {
		rowList := []int{}
		for _, letter := range row {
			val, _ := strconv.Atoi(string(letter))
			rowList = append(rowList, val)
		}
		inputMap = append(inputMap, rowList)
	}
	fmt.Println(part1(inputMap))
	fmt.Println(part2(inputMap))
}