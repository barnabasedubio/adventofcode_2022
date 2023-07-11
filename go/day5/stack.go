package main

import "fmt"

type Stack struct {
	elements []string
}

func (s *Stack) Push(value string) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() (string, error) {
	if len(s.elements) == 0 {
		return "", fmt.Errorf("Can not pop off empty stack")
	}
	popped := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return popped, nil
}

func (s *Stack) Reverse() {
	for i, j := 0, len(s.elements)-1; i < j; i, j = i+1, j-1 {
		s.elements[i], s.elements[j] = s.elements[j], s.elements[i]
	}
}
