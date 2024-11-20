package _4

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

//go:embed input.txt
var input string

//go:embed sample.txt
var sample string

func TestSolution(t *testing.T) {
	type test struct {
		name            string
		input           string
		expectedPartOne int64
		expectedPartTwo int64
	}

	tests := []test{
		{
			name:            "with sample",
			input:           sample,
			expectedPartOne: 35,
			expectedPartTwo: 46,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			seeds, conversionMaps := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(seeds, conversionMaps); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(seeds, conversionMaps); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})

	}

}

type ConversionMapRange struct {
	destinationRangeStart int64
	sourceRangeStart      int64
	rangeLen              int64
}

func convertStrListToInt(strList string) []int64 {
	nums := make([]int64, 0)
	for _, numStr := range strings.Split(strList, " ") {
		num, err := strconv.ParseInt(strings.TrimSpace(numStr), 10, 64)
		if err == nil {
			nums = append(nums, num)
		}

	}
	return nums
}

func parseInput(input string) ([]int64, [][]ConversionMapRange) {
	lines := strings.Split(input, "\n")

	seeds := convertStrListToInt(strings.Replace(lines[0], "seeds: ", "", 1))

	conversionMaps := make([][]ConversionMapRange, 0)
	conversionMapRange := make([]ConversionMapRange, 0)

	conversionRegex := regexp.MustCompile(`^(?P<drs>\d+) (?P<srs>\d+) (?P<rl>\d+)$`)

	for i := 1; i < len(lines); i++ {
		if strings.Contains(lines[i], "map") {
			if len(conversionMapRange) > 0 {
				conversionMaps = append(conversionMaps, conversionMapRange)
				conversionMapRange = make([]ConversionMapRange, 0)
			}
		} else {
			matches := conversionRegex.FindStringSubmatch(lines[i])
			if len(matches) > 3 {
				drs, _ := strconv.ParseInt(matches[1], 10, 64)
				srs, _ := strconv.ParseInt(matches[2], 10, 64)
				rl, _ := strconv.ParseInt(matches[3], 10, 64)
				conversionMapRange = append(conversionMapRange, ConversionMapRange{drs, srs, rl})
			}
		}
	}

	conversionMaps = append(conversionMaps, conversionMapRange)

	return seeds, conversionMaps
}

func applyConversionMapsToNum(num int64, curConv int, conversionMaps [][]ConversionMapRange) int64 {
	convertedNum := num
	for _, conversionMap := range conversionMaps[curConv] {
		if conversionMap.sourceRangeStart <= num && num <= conversionMap.sourceRangeStart+conversionMap.rangeLen-1 {
			convertedNum = conversionMap.destinationRangeStart + (num - conversionMap.sourceRangeStart)
			break
		}
	}

	if curConv == len(conversionMaps)-1 {
		return convertedNum
	}

	return applyConversionMapsToNum(convertedNum, curConv+1, conversionMaps)

}

func partOne(seeds []int64, conversionMaps [][]ConversionMapRange) int64 {
	lowestLoc := int64(999999999999999999)

	for _, seed := range seeds {
		loc := applyConversionMapsToNum(seed, 0, conversionMaps)
		if loc < lowestLoc {
			lowestLoc = loc
		}
	}
	return lowestLoc
}

type numTuple struct {
	from int64
	to   int64
}

func convert(num int64, conversion ConversionMapRange) int64 {
	return conversion.sourceRangeStart - conversion.destinationRangeStart + num

}

func applyConversionToNumTuples(tuples []numTuple, curConv int, conversionMap [][]ConversionMapRange) []numTuple {
	newTuples := make([]numTuple, 0)

	for i := 0; i < len(tuples); i++ {
		num := tuples[i]
		intersect := false
		for _, conversion := range conversionMap[curConv] {
			delta := conversion.destinationRangeStart - conversion.sourceRangeStart
			source := numTuple{conversion.sourceRangeStart, conversion.sourceRangeStart + conversion.rangeLen - 1}

			// case 1: source is completely in tuple range sf nf nt st
			if source.from <= num.from && num.to <= source.to {
				intersect = true
				newTuples = append(newTuples, numTuple{num.from + delta, num.to + delta})
				// case 2: source is partially in tuple range sf nf st nt
			} else if source.from <= num.from && num.from <= source.to {
				intersect = true
				newTuples = append(newTuples, numTuple{num.from + delta, source.to + delta})
				tuples = append(tuples, numTuple{source.to + 1, num.to})

				// case 3: source is partially in tuple range nf sf nt st
			} else if num.from <= source.from && source.from <= num.to {
				intersect = true
				newTuples = append(newTuples, numTuple{source.from + delta, num.to + delta})
				tuples = append(tuples, numTuple{num.from, source.from - 1})
			}

			if intersect {
				break
			}
		}
		if !intersect {
			newTuples = append(newTuples, num)
		}

	}

	if curConv == len(conversionMap)-1 {
		return newTuples
	}

	return applyConversionToNumTuples(newTuples, curConv+1, conversionMap)
}

func partTwo(seeds []int64, conversionMaps [][]ConversionMapRange) int64 {
	seedTuples := make([]numTuple, 0)
	for i := 0; i < len(seeds)-1; i += 2 {
		seedTuples = append(seedTuples, numTuple{seeds[i], seeds[i] + seeds[i+1] - 1})
	}

	tuplesAfterConversion := applyConversionToNumTuples(seedTuples, 0, conversionMaps)
	lowestLoc := int64(999999999999999999)
	for _, tuple := range tuplesAfterConversion {
		if tuple.from <= lowestLoc {
			lowestLoc = tuple.from
		}
	}

	return lowestLoc
}
