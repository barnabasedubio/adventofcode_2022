package main

import (
	"log"
	"fmt"
)

func GetHeight(letter string) int {
	var letterAsRune rune
	switch letter {
	case "a", "S":
		letterAsRune = []rune("a")[0]
	case "z", "E":
		letterAsRune = []rune("z")[0]
	default:
		letterAsRune = []rune(letter)[0]
	}
	return int(letterAsRune) - 96
}

func GetNeighbours(current Node, nodeList *[]Node, inputMatrix [][]string) []*Node {
	neighbours := []*Node{}
	// top
	if current.Row-1 >= 0 {
		for _, elem := range *nodeList {
			if (elem.Row == current.Row-1 && elem.Col == current.Col) &&
				(elem.Height-current.Height <= 1) {
				top := &elem
				neighbours = append(neighbours, top)
				break
			}
		}
	}
	// bottom
	if current.Row+1 <= len(inputMatrix) {
		for _, elem := range *nodeList {
			if (elem.Row == current.Row+1 && elem.Col == current.Col) &&
				(elem.Height-current.Height <= 1) {
				bottom := &elem
				neighbours = append(neighbours, bottom)
				break
			}
		}
	}
	// left
	if current.Col-1 >= 0 {
		for _, elem := range *nodeList {
			if (elem.Col == current.Col-1 && elem.Row == current.Row) &&
				(elem.Height-current.Height <= 1) {
				left := &elem
				neighbours = append(neighbours, left)
				break
			}
		}
	}
	// right
	if current.Col+1 <= len(inputMatrix[0]) {
		for _, elem := range *nodeList {
			if (elem.Col == current.Col+1 && elem.Row == current.Row) &&
				(elem.Height-current.Height <= 1) {
				right := &elem
				neighbours = append(neighbours, right)
				break
			}
		}
	}
	fmt.Println()
	for i := range neighbours {
		log.Println(*neighbours[i])
	}
	return neighbours
}
