package main

import (
	"fmt"
	"os"
	"strings"
)

var itemValue = map[string]int{
	"X": 1, // rock
	"Y": 2, // paper
	"Z": 3, // scissors
}

var outcomeValues = map[string]map[string]int{
	"A": { // opponent plays rock
		"X": 3,
		"Y": 6,
		"Z": 0,
	},
	"B": { // opponent plays paper
		"X": 0,
		"Y": 3,
		"Z": 6,
	},
	"C": { // opponent plays scissors
		"X": 6,
		"Y": 0,
		"Z": 3,
	},
}

// required for part 2
var results = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

func part1(inputList []string) int {
	total := 0
	for _, val := range inputList {
		opp := strings.Split(val, " ")[0]
		mine := strings.Split(val, " ")[1]
		total += itemValue[mine] + outcomeValues[opp][mine]
	}
	return total
}

func part2(inputList []string) int {
	total := 0
	for _, val := range inputList {
		opp := strings.Split(val, " ")[0]
		outcome := strings.Split(val, " ")[1] // X: loss, Y: draw, Z: win
		result := results[outcome]
		// iterate over outcomeValues of the opponenes hand and select the key
		// corresponding to the outcome of result
		for key, value := range outcomeValues[opp] {
			if value == result {
				total += itemValue[key] + outcomeValues[opp][key]
			}
		}
	}
	return total
}

func main() {
	input, _ := os.ReadFile("day2.in")
	inputList := strings.Split(string(input), "\n")
	fmt.Println(part1(inputList))
	fmt.Println(part2(inputList))
}
