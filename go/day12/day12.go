package main

import (
	"fmt"
	"math"
	"os"

	//"slices"
	"strings"
)

type Node struct {
	Letter     string
	Height     int
	Row        int
	Col        int
	Distance   int
	Visited    bool
	Neighbours []*Node
	Previous   *Node
}

func part1(inputMatrix [][]string) int {
	// we'll be using Dijstra's algorithm to get the shortest path between
	// the source and the destination node
	// a given node x has neighbor y only if y's height is at most 1 higher than that of x.
	// ==> y.Height - x.Height <= 1
	nodes := []Node{}
	queue := []*Node{}
	for i := range inputMatrix {
		for j := range inputMatrix[i] {
			letter := inputMatrix[i][j]
			distance := math.MaxInt32
			if letter == "S" {
				distance = 0
			}
			newNode := Node{
				Letter:     letter,
				Height:     GetHeight(letter),
				Row:        i,
				Col:        j,
				Distance:   distance,
				Visited:    false,
				Neighbours: nil,
				Previous:   nil,
			}
			nodes = append(nodes, newNode)
		}
	}
	for i := range nodes {
		nodes[i].Neighbours = GetNeighbours(nodes[i], nodes, inputMatrix)
		queue = append(queue, &nodes[i])
	}

	toBeVisited := len(queue)

	for toBeVisited > 0 {
		// retrieve node with shortest distance from source
		minDist := math.MaxInt32
		minDistIndex := -1
		for i := range queue {
			if !queue[i].Visited && queue[i].Distance <= minDist {
				minDist = queue[i].Distance
				minDistIndex = i
			}
		}
		for i := range queue[minDistIndex].Neighbours {
			if !queue[minDistIndex].Neighbours[i].Visited {
				if queue[minDistIndex].Distance+1 < queue[minDistIndex].Neighbours[i].Distance {
					queue[minDistIndex].Neighbours[i].Distance = queue[minDistIndex].Distance + 1
					queue[minDistIndex].Neighbours[i].Previous = queue[minDistIndex]
				}
			}
		}
		// remove node from queue
		// queue = append(queue[:minDistIndex], queue[minDistIndex+1:]...)
		queue[minDistIndex].Visited = true
		toBeVisited--
	}
	fmt.Println(nodes)

	return -1
}

func main() {
	input, _ := os.ReadFile("day12.test.in")
	inputList := strings.Split(string(input), "\n")
	inputMatrix := make([][]string, len(inputList))
	for i := range inputMatrix {
		inputMatrix[i] = strings.Split(inputList[i], "")
	}
	fmt.Println(part1(inputMatrix))
}
