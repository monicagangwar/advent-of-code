package main

import (
	"fmt"
	"sort"
)

func main() {
	monkeys := getMonkeysFromInput()
	compute(monkeys, 20, true)
	monkeys = getMonkeysFromInput()
	compute(monkeys, 10000, false)

}

func compute(monkeys []*monkey, rounds int, divideWorry bool) {
	var mod int
	if !divideWorry {
		mod = 1
		for _, monkey := range monkeys {
			mod *= monkey.test[0]
		}
	}
	inspectedItems := make([]int64, len(monkeys))
	for round := 1; round <= rounds; round++ {
		for monkeyIdx, monkey := range monkeys {
			ptr := monkey.itemsHead
			for {
				if ptr == nil {
					break
				}
				inspectedItems[monkeyIdx]++
				newItem := monkey.performOp(ptr.val)
				if divideWorry {
					newItem /= 3
				} else {
					newItem %= mod
				}
				throwToMonkey := monkey.performTest(newItem)
				monkeys[throwToMonkey].addItem(newItem)
				monkey.removeItemFromStart()
				ptr = ptr.next
			}
		}
		//display(round, monkeys)
	}
	sort.Slice(inspectedItems, func(i, j int) bool {
		return inspectedItems[i] > inspectedItems[j]
	})
	fmt.Println(inspectedItems)
	fmt.Printf("%d\n", inspectedItems[0]*inspectedItems[1])
}

func display(round int, monkeys []*monkey) {
	fmt.Printf("\n============ Round %d ============", round)

	for idx, monkey := range monkeys {
		fmt.Printf("\nMonkey %d:", idx)
		monkey.displayItems()
	}
}
