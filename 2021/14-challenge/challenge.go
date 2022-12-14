package main

import (
	"bufio"
	"fmt"
	"runtime"
	"strings"

	"github.com/monicagangwar/advent-of-code/input"
)

func main() {
	challenge(10)
	challenge(40)
}

func challenge(steps int) {
	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	recordMap := make(map[string]int64, 0)
	substitutionMap := make(map[string]string, 0)
	charCount := make(map[string]int64, 0)

	inputLineScanned := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		if !inputLineScanned {
			for idx := 0; idx < len(line); idx++ {
				if idx < len(line)-1 {
					pattern := line[idx : idx+2]
					if _, found := recordMap[pattern]; !found {
						recordMap[pattern] = 0
					}
					recordMap[pattern] += 1
				}
				if _, found := charCount[string(line[idx])]; !found {
					charCount[string(line[idx])] = 0
				}
				charCount[string(line[idx])] += 1
			}
			inputLineScanned = true
		} else {
			substitutionRecord := strings.Split(line, " -> ")
			substitutionMap[substitutionRecord[0]] = substitutionRecord[1]
		}
	}

	for step := 1; step <= steps; step++ {
		newRecordMap := make(map[string]int64, 0)
		for pattern, patternCount := range recordMap {

			substitue := substitutionMap[pattern]

			leftPattern := fmt.Sprintf("%s%s", string(pattern[0]), substitue)
			rightPattern := fmt.Sprintf("%s%s", substitue, string(pattern[1]))

			if _, found := charCount[substitutionMap[pattern]]; !found {
				charCount[substitutionMap[pattern]] = 0
			}

			charCount[substitutionMap[pattern]] += patternCount

			if _, found := newRecordMap[leftPattern]; !found {
				newRecordMap[leftPattern] = 0
			}
			newRecordMap[leftPattern] += patternCount

			if _, found := newRecordMap[rightPattern]; !found {
				newRecordMap[rightPattern] = 0
			}
			newRecordMap[rightPattern] += patternCount
		}
		recordMap = newRecordMap
		//fmt.Printf("Step: %d, charCount: %+v\n", step, charCount)
	}

	mostCommon := int64(0)

	for _, count := range charCount {
		if count > mostCommon {
			mostCommon = count
		}
	}

	leastCommon := mostCommon
	for _, count := range charCount {
		if count < leastCommon {
			leastCommon = count
		}
	}

	fmt.Printf("%d\n", mostCommon-leastCommon)
}
