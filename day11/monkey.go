package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/monicagangwar/advent-of-code-2022/input"
)

func getMonkeysFromInput() []*monkey {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")
	monkeys := make([]*monkey, 0)

	idx := 1
	for {
		if idx >= len(lines) {
			break
		}
		monkey := getMonkeyFromInput(lines[idx : idx+5])
		idx += 7
		monkeys = append(monkeys, monkey)
	}
	return monkeys
}

func getMonkeyFromInput(lines []string) *monkey {
	m := &monkey{
		operation: make([]string, 0),
	}

	// get items
	lines[0] = strings.ReplaceAll(lines[0], "Starting items: ", "")
	lines[0] = strings.ReplaceAll(lines[0], " ", "")
	nums := strings.Split(lines[0], ",")
	for _, numStr := range nums {
		num, _ := strconv.Atoi(numStr)
		m.addItem(num)
	}

	// get operation
	lines[1] = strings.TrimSpace(lines[1])
	lines[1] = strings.ReplaceAll(lines[1], "Operation: new = ", "")
	m.operation = strings.Split(lines[1], " ")

	// get test
	m.test[0] = getTest(lines[2], "Test: divisible by")
	m.test[1] = getTest(lines[3], "If true: throw to monkey")
	m.test[2] = getTest(lines[4], "If false: throw to monkey")

	return m
}

func getTest(line string, replaceStr string) int {
	line = strings.ReplaceAll(line, replaceStr, "")
	line = strings.ReplaceAll(line, " ", "")
	num, _ := strconv.Atoi(line)
	return num
}

type itemNode struct {
	val  int
	next *itemNode
}

type monkey struct {
	itemsHead *itemNode
	itemsTail *itemNode
	operation []string
	test      [3]int
}

func (m *monkey) performOp(item int) int {
	var lop, rop int
	if m.operation[0] == "old" {
		lop = item
	} else {
		lop, _ = strconv.Atoi(m.operation[0])
	}
	if m.operation[2] == "old" {
		rop = item
	} else {
		rop, _ = strconv.Atoi(m.operation[2])
	}

	switch m.operation[1] {
	case "+":
		return lop + rop
	case "*":
		return lop * rop
	}

	panic(fmt.Sprintf("op %s not handled", m.operation[1]))
}

func (m *monkey) performTest(item int) int {
	if item%m.test[0] == 0 {
		return m.test[1]
	}
	return m.test[2]
}

func (m *monkey) addItem(x int) {
	curItemNode := itemNode{
		val:  x,
		next: nil,
	}
	if m.itemsHead == nil {
		m.itemsHead = &curItemNode
	}
	if m.itemsTail == nil {
		m.itemsTail = &curItemNode
	} else {
		m.itemsTail.next = &curItemNode
		m.itemsTail = &curItemNode
	}
}

func (m *monkey) removeItemFromStart() {
	if m.itemsHead != nil {
		m.itemsHead = m.itemsHead.next
		if m.itemsHead == nil {
			m.itemsTail = nil
		}
	}
}

func (m *monkey) displayItems() {
	ptr := m.itemsHead
	for {
		if ptr == nil {
			break
		}
		fmt.Printf("%d ", ptr.val)
		ptr = ptr.next
	}
}
