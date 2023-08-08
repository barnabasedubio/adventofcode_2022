package main

import (
	"aoc2022/shared"
	"fmt"
)

type Monkey struct {
	Id             int
	Items          shared.Queue[int64]
	CurrentItem    int64
	Operation      string
	Operand        int64
	DivisibleBy    int64
	TrueCase       int
	FalseCase      int
	TimesInspected int64
}

func (m *Monkey) HasItems() bool {
	return !m.Items.IsEmpty()
}

func (m *Monkey) Inspect() {
	m.CurrentItem = m.Items.Pop()
	// fmt.Printf("Monkey %v inspecting item %v. Items in queue: %v\n", m.Id, m.CurrentItem, m.Items)
	switch m.Operation {
	case "add":
		m.CurrentItem += m.Operand
	case "mult":
		m.CurrentItem *= m.Operand
	case "square":
		m.CurrentItem *= m.CurrentItem
	}
	m.TimesInspected++
}

func (m *Monkey) ReduceWorryLevel() {
	m.CurrentItem /= 3
}

func (m *Monkey) Test() int {
	if m.CurrentItem%m.DivisibleBy == 0 {
		return m.TrueCase
	} else {
		return m.FalseCase
	}
}

func (m *Monkey) Throw(t *Monkey) {
	t.Receive(m.CurrentItem)
}

func (m *Monkey) Receive(val int64) {
	m.Items.Push(val)
}

func (m Monkey) String() string {
	return fmt.Sprintf("Items: %v\nTimes inspected: %v\n", m.Items, m.TimesInspected)
}
