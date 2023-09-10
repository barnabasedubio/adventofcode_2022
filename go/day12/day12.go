package main

import (
	"fmt"
	"log"
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
	nodes := []Node{}
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
		nodes[i].Neighbours = GetNeighbours(nodes[i], &nodes, inputMatrix)
	}

	toBeVisited := len(nodes)

	for toBeVisited > 0 {
		// perform Dijkstra's algorithm
		// retrieve node with shortest distance from source
		minDist := math.MaxInt32
		minDistIndex := -1
		for i := range nodes {
			if !nodes[i].Visited && nodes[i].Distance <= minDist {
				minDist = nodes[i].Distance
				minDistIndex = i
			}
		}
		fmt.Printf("\n\n-------------------------\n\n")
		log.Printf("Looking at neighbors of %v\n", nodes[minDistIndex])
		for i := range nodes[minDistIndex].Neighbours {
			log.Printf("found %v\n", nodes[minDistIndex].Neighbours[i])
			if !nodes[minDistIndex].Neighbours[i].Visited {
				if nodes[minDistIndex].Distance+1 < nodes[minDistIndex].Neighbours[i].Distance {
					log.Println("update distance...")
					log.Println(nodes[minDistIndex].Neighbours[i])
					nodes[minDistIndex].Neighbours[i].Distance = nodes[minDistIndex].Distance + 1
					nodes[minDistIndex].Neighbours[i].Previous = &nodes[minDistIndex]
					log.Println(nodes[minDistIndex].Neighbours[i])
				}
			} else {
				log.Println("Has already been visited. Skipping...")
			}
		}
		nodes[minDistIndex].Visited = true
		toBeVisited--
	}
	// log.Println(nodes)

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
