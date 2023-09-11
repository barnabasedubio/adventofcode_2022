package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
	"time"
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

func dijkstra(inputMatrix [][]string, sourceRow int, sourceCol int, done chan bool, result chan int) {
	defer func() {
		done <- true
	}()

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
			result <- node.Distance
			break
		}
	}
}

func part1(inputMatrix [][]string) int {
	startTime := time.Now()
	done := make(chan bool)
	result := make(chan int)
	for i := range inputMatrix {
		for j := range inputMatrix[i] {
			if inputMatrix[i][j] == "S" {
				go dijkstra(inputMatrix, i, j, done, result)
			}
		}
	}
	minDistance := <-result
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Completed part 1 in %v seconds\n", elapsedTime.Seconds())
	return minDistance

}

func part2(inputMatrix [][]string) int {
	startTime := time.Now()
	numWorkers := 0
	done := make(chan bool)
	results := make(chan int)
	nodesWithHeight1 := []int{}
	for i := range inputMatrix {
		for j := range inputMatrix[i] {
			if inputMatrix[i][j] == "a" || inputMatrix[i][j] == "S" {
				numWorkers++
				go dijkstra(inputMatrix, i, j, done, results)
			}
		}
	}
	for i := 0; i < numWorkers; i++ {
		nodesWithHeight1 = append(nodesWithHeight1, <-results)
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Completed part 2 in %v seconds\n", elapsedTime.Seconds())
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
