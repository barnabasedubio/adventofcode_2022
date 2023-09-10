package main

import (
	"fmt"
	"math"
	"os"
	"slices"
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

func dijkstra(inputMatrix [][]string, sourceRow int, sourceCol int) int {
	nodes := []Node{}
	for i := range inputMatrix {
		for j := range inputMatrix[i] {
			letter := inputMatrix[i][j]
			distance := math.MaxInt32
			if sourceRow == i && sourceCol == j {
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
		SetNeighbours(&nodes[i], &nodes, inputMatrix)
	}

	toBeVisited := len(nodes)

	for toBeVisited > 0 {
		// perform Dijkstra's algorithm

		minDist := math.MaxInt32
		minDistIndex := -1
		for i := range nodes {
			if !nodes[i].Visited && nodes[i].Distance <= minDist {
				minDist = nodes[i].Distance
				minDistIndex = i
			}
		}
		for i := range nodes[minDistIndex].Neighbours {
			if !nodes[minDistIndex].Neighbours[i].Visited {
				if nodes[minDistIndex].Distance+1 < nodes[minDistIndex].Neighbours[i].Distance {
					nodes[minDistIndex].Neighbours[i].Distance = nodes[minDistIndex].Distance + 1
					nodes[minDistIndex].Neighbours[i].Previous = &nodes[minDistIndex]
				}
			}
		}
		nodes[minDistIndex].Visited = true
		toBeVisited--
	}

	for _, node := range nodes {
		if node.Letter == "E" {
			return node.Distance
		}
	}
	return -1
}

func part1(inputMatrix [][]string) int {
	for i := range inputMatrix {
		for j := range inputMatrix[i] {
			if inputMatrix[i][j] == "S" {
				return dijkstra(inputMatrix, i, j)
			}
		}
	}
	return -1
}

func part2(inputMatrix [][]string) int {
	fmt.Println("Solving part 2. This will take a while...")
	nodesWithHeight1 := []int{}
	for i := range inputMatrix {
		for j := range inputMatrix[i] {
			if inputMatrix[i][j] == "a" || inputMatrix[i][j] == "S" {
				shortestPath := dijkstra(inputMatrix, i, j)
				nodesWithHeight1 = append(nodesWithHeight1, shortestPath)
			}
		}
	}
	return slices.Min(nodesWithHeight1)
}

func main() {
	input, _ := os.ReadFile("day12.in")
	inputList := strings.Split(string(input), "\n")
	inputMatrix := make([][]string, len(inputList))
	for i := range inputMatrix {
		inputMatrix[i] = strings.Split(inputList[i], "")
	}
	fmt.Println(part1(inputMatrix))
	fmt.Println(part2(inputMatrix))
}
