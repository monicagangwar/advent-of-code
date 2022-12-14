package main

import (
	"fmt"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/monicagangwar/advent-of-code/input"
)

func main() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")
	rightOrderIndicesSum := 0
	packets := make([]Item, 0)
	for idx := 0; idx < len(lines)-1; idx += 3 {
		packet1, packet2 := convertStrToItem(lines[idx]), convertStrToItem(lines[idx+1])
		packets = append(packets, packet1, packet2)
		if isRightOrder(packet1, packet2) == 1 {
			rightOrderIndicesSum += (idx / 3) + 1
		}
	}
	fmt.Println(rightOrderIndicesSum)
	packets = append(packets, convertStrToItem("[[2]]"), convertStrToItem("[[6]]"))
	sort.Slice(packets, func(i, j int) bool {
		return isRightOrder(packets[i], packets[j]) >= 0
	})
	decoderKey := 1
	for idx, packet := range packets {
		packetStr := packet.String()
		if packetStr == "[[2]]" || packetStr == "[[6]]" {
			decoderKey *= idx + 1
		}
	}
	fmt.Println(decoderKey)
}

func isRightOrder(left Item, right Item) int {
	if left.valInt != nil && right.valInt != nil {
		return del(*left.valInt, *right.valInt)
	}
	if left.valList != nil && right.valList != nil {
		leftIdx, rightIdx := len(left.valList)-1, len(right.valList)-1
		for {
			if leftIdx < 0 || rightIdx < 0 {
				break
			}
			rightOrder := isRightOrder(left.valList[leftIdx], right.valList[rightIdx])
			if rightOrder != 0 {
				return rightOrder
			}
			leftIdx--
			rightIdx--

		}
		return del(len(left.valList), len(right.valList))
	}
	// convert left to list and compare
	if left.valInt != nil {
		return isRightOrder(Item{
			valList: []Item{{valInt: left.valInt}},
		}, right)
	}
	// convert right to list and compare
	return isRightOrder(left, Item{
		valList: []Item{{valInt: right.valInt}},
	})
}

func del(leftOp int, rightOp int) int {
	delta := rightOp - leftOp
	if delta == 0 {
		return 0
	}
	if delta > 0 {
		return 1
	}
	return -1
}

type Item struct {
	valInt  *int
	valList []Item
}

func (i *Item) String() string {
	if i.valInt != nil {
		return fmt.Sprintf("%d", *i.valInt)
	}
	if i.valList != nil {
		str := "["
		for _, item := range i.valList {
			str += item.String() + ","
		}
		if str[len(str)-1] == ',' {
			str = str[:len(str)-1]
		}
		str += "]"
		return str
	}
	return ""
}

func convertStrToItem(line string) Item {
	stack := make([]interface{}, 0)
	var num *int
	for strIdx := 0; strIdx < len(line); {
		switch line[strIdx] {
		case '[':
			stack = append(stack, '[')
			break
		case ',':
			if num != nil {
				stack = append(stack, Item{valInt: num})
				num = nil
			}
			break
		case ']':
			if num != nil {
				stack = append(stack, Item{valInt: num})
				num = nil
			}
			idx := len(stack) - 1
			newItem := Item{
				valList: make([]Item, 0),
			}
			for {
				if stack[idx] == '[' {
					stack = stack[:idx]
					stack = append(stack, newItem)
					break
				}
				newItem.valList = append(newItem.valList, stack[idx].(Item))
				idx--
			}
			break
		default:
			digit, _ := strconv.Atoi(string(line[strIdx]))
			if num != nil {
				digit = (*num * 10) + digit
			}
			num = &digit
			break
		}
		strIdx++
	}
	return stack[0].(Item)
}
