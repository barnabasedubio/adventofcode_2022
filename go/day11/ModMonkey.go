package main

import (
	"aoc2022/shared"
	"fmt"
)

/*
Example:
Monkey 1 contains items [12, 34, 456 4] and is divisible by 23
Monkey 2 contains items [44, 2, 54, 1] and is divisible by 12

Monkey 1 item list will be: [
	[12 % 23, 12 % 12],
	[34 % 23, 34 % 12],
	[456 % 23, 456 % 12],
	[4 % 23, 4 % 12]
]
Monkey 2 item list will be: [
	[44 % 23, 44 % 12],
	[2 % 23, 2 % 12],
	[54 % 23, 54 % 12],
	[1 % 23, 1 % 12]
]
*/

type ModMonkey struct {
	Id             int
	Items          shared.Queue[[]int64]
	CurrentItem    []int64
	Operation      string
	Operand        int64
	Divisors       []int64
	TrueCase       int
	FalseCase      int
	TimesInspected int64
}

func (m *ModMonkey) HasItems() bool {
	return !m.Items.IsEmpty()
}

func (m *ModMonkey) Inspect() {
	m.CurrentItem = m.Items.Pop()
	switch m.Operation {
	case "add":
		for i := range m.CurrentItem {
			m.CurrentItem[i] = (m.CurrentItem[i] + m.Operand) % m.Divisors[i]
		}
	case "mult":
		for i := range m.CurrentItem {
			m.CurrentItem[i] = (m.CurrentItem[i] * m.Operand) % m.Divisors[i]
		}
	case "square":
		for i := range m.CurrentItem {
			m.CurrentItem[i] = (m.CurrentItem[i] * m.CurrentItem[i]) % m.Divisors[i]
		}
	}
	m.TimesInspected++
}

func (m *ModMonkey) Test() int {
	if m.CurrentItem[m.Id]%m.Divisors[m.Id] == 0 { // equivalent to m.CurrentItem[m.Id] == 0
		return m.TrueCase
	} else {
		return m.FalseCase
	}
}

func (m *ModMonkey) Throw(t *ModMonkey) {
	t.Receive(m.CurrentItem)
}

func (m *ModMonkey) Receive(val []int64) {
	m.Items.Push(val)
}
func (m ModMonkey) String() string {
	return fmt.Sprintf("Times inspected: %v\n", m.TimesInspected)
}
