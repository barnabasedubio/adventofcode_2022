package main

import (
	"fmt"
	"os"
	"strings"
)

func isSet(letters []rune) bool {
	counter := 0
	charSet := map[rune]bool{}
	for _, letter := range letters {
		if !charSet[letter] {
			charSet[letter] = true
			counter++
		}
	}
	return counter == len(letters)
}

func part1(sequence string, part int) int {
	letters := []rune{}
	for i, letter := range sequence {
		letters = append(letters, letter)
		switch part {
		case 1:
			if i > 2 {
				if isSet(letters[len(letters)-4:]) {
					return i + 1
				}
			}
		case 2:
			if i > 12 {
				if isSet(letters[len(letters)-14:]) {
					return i + 1
				}
			}
		}
	}
	return -1
}

func main() {
	input, _ := os.ReadFile("day6.in")
	inputList := strings.Split(string(input), "\n")
	for _, sequence := range inputList {
		fmt.Println(part1(sequence, 1))
		fmt.Println(part1(sequence, 2))
	}
}
