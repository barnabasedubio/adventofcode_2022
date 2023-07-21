package shared

import "fmt"

type Stack[T any] struct {
	Elements []T
}

func (s *Stack[T]) Push(value T) {
	s.Elements = append(s.Elements, value)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.Elements) == 0 {
		var zero T
		return zero, fmt.Errorf("Can not pop off empty stack")
	}
	popped := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return popped, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.Elements) == 0
}

func (s *Stack[T]) Reverse() {
	for i, j := 0, len(s.Elements)-1; i < j; i, j = i+1, j-1 {
		s.Elements[i], s.Elements[j] = s.Elements[j], s.Elements[i]
	}
}
