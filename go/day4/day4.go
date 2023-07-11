package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(inputList []string, part int) int {
	total := 0
	for _, line := range inputList {
		sec1 := strings.Split(line, ",")[0]
		sec2 := strings.Split(line, ",")[1]

		sec1begin, _ := strconv.Atoi(strings.Split(sec1, "-")[0])
		sec1end, _ := strconv.Atoi(strings.Split(sec1, "-")[1])

		sec2begin, _ := strconv.Atoi(strings.Split(sec2, "-")[0])
		sec2end, _ := strconv.Atoi(strings.Split(sec2, "-")[1])

		switch part {
		case 1:
			if (sec1begin <= sec2begin && sec1end >= sec2end) ||
				(sec2begin <= sec1begin && sec2end >= sec1end) {
				total += 1
			}
		case 2:
			if (sec1begin <= sec2begin && sec2begin <= sec1end) ||
				(sec2begin <= sec1begin && sec1begin <= sec2end) {
				total += 1
			}
		}
	}
	return total
}

func main() {
	input, _ := os.ReadFile("day4.in")
	inputList := strings.Split(string(input), "\n")
	fmt.Println(solve(inputList, 1))
	fmt.Println(solve(inputList, 2))
}
