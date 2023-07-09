package main

import (
	"fmt"
	"os"
	"strings"
)

func getPriority(letter rune) int {
	if int(letter) >= 97 && int(letter) <= 122 {
		return int(letter) - 96
	}
	if int(letter) >= 65 && int(letter) <= 90 {
		return int(letter) - 38
	}
	return -1
}

func makeSet(s string) []rune {
	charSet := map[rune]bool{}
	set := []rune{}
	for _, letter := range s {
		if !charSet[letter] {
			charSet[letter] = true
			set = append(set, letter)
		}
	}
	return set
}

func populateCharSet(m *map[rune]int, s []rune) {
	for _, char := range s {
		_, ok := (*m)[char]
		if !ok {
			(*m)[char] = 1
		} else {
			(*m)[char] += 1
		}
	}
}

func part1(inputList []string) int {
	total := 0
	for _, line := range inputList {
		charSet := map[rune]int{}
		left := makeSet(line[:len(line)/2])
		right := makeSet(line[len(line)/2:])
		populateCharSet(&charSet, left)
		populateCharSet(&charSet, right)
		for char, amount := range charSet {
			if amount == 2 {
				total += getPriority(char)
			}
		}

	}
	return total
}

func part2(inputList []string) int {
	total := 0
	for i := 0; i < len(inputList)-2; i += 3 {
		letterMap := map[rune]int{}
		s1 := makeSet(inputList[i])
		s2 := makeSet(inputList[i+1])
		s3 := makeSet(inputList[i+2])
		populateCharSet(&letterMap, s1)
		populateCharSet(&letterMap, s2)
		populateCharSet(&letterMap, s3)
		for letter, amount := range letterMap {
			if amount == 3 {
				total += getPriority(letter)
			}
		}
	}
	return total
}

func main() {
	input, _ := os.ReadFile("day3.in")
	inputList := strings.Split(string(input), "\n")
	fmt.Println(part1(inputList))
	fmt.Println(part2(inputList))
}
