package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func Solve(inputList []string, part int) (int, error) {
	current := 0
	var listOfCapacities []int

	for _, val := range inputList {
		val = strings.Trim(val, "\n")
		if val == "" {
			listOfCapacities = append(listOfCapacities, current)
			current = 0
		} else {
			if valAsInt, err := strconv.Atoi(val); err != nil {
				return -1, fmt.Errorf("Error: unable to convert value to int: %q\n", val)
			} else {
				current += valAsInt
			}
		}
	}

	sort.Ints(listOfCapacities)

	switch part {
	case 1:
		return listOfCapacities[len(listOfCapacities)-1], nil
	case 2:
		return Sum(listOfCapacities[len(listOfCapacities)-3:]), nil
	default:
		return -1, fmt.Errorf("Unknown parameter part: %v", part)
	}
}

func main() {
	var inputString string
	if input, err := os.ReadFile("day1.in"); err != nil {
		panic("There was a problem reading the input file.")
	} else {
		inputString = string(input)
	}
	inputList := strings.Split(inputString, "\n")

	part := 1
	if res, err := Solve(inputList, part); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
