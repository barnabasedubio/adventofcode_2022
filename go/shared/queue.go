package shared

type Queue[T any] struct {
	Elements []T
}

func (q *Queue[T]) Push(value T) {
	q.Elements = append(q.Elements, value)
}

func (q *Queue[T]) Pop() T {
	if len(q.Elements) == 0 {
		panic("Can not pop off empty Queue!")
	}
	popped := q.Elements[0]
	if len(q.Elements) == 1 {
		q.Elements = []T{}
	} else {
		q.Elements = q.Elements[1:]
	}
	return popped
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q *Queue[T]) Reverse() {
	for i, j := 0, len(q.Elements)-1; i < j; i, j = i+1, j-1 {
		q.Elements[i], q.Elements[j] = q.Elements[j], q.Elements[i]
	}
}
