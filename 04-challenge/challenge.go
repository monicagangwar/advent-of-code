package main

import (
	"bufio"
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/monicagangwar/advent-of-code-2021/input"
)

func main() {
	//one()
	two()
}

func one() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	calledNumbers := make([]int, 120)
	calledNumbersScanned := false
	bingoBoards := make([][][]int, 0)
	bingoBoard := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if !calledNumbersScanned {
			calledNumbersScanned = true
			for idx, strNum := range strings.Split(line, ",") {
				num, _ := strconv.Atoi(strNum)
				calledNumbers[num] = idx + 1
			}
		}

		if strings.Trim(line, " ") == "" {
			bingoBoard = make([][]int, 0)
		} else {
			row := make([]int, 0)
			for _, strNum := range strings.Split(line, " ") {
				num, err := strconv.Atoi(strNum)
				if err != nil {
					continue
				}
				row = append(row, num)
			}
			bingoBoard = append(bingoBoard, row)
			if len(bingoBoard) == 5 {
				bingoBoards = append(bingoBoards, bingoBoard)
			}
		}
	}

	winnerTicker := 110
	winnerNum := -1
	winnerBoardIdx := -1

	for boardIdx, board := range bingoBoards {

		for row := 0; row < 5; row++ {
			rowMarked := true
			colMarked := true
			localRowNum := -1
			localRowTicker := -1
			localColNum := -1
			localColTicker := -1
			for col := 0; col < 5; col++ {
				rowTicker := calledNumbers[board[row][col]]
				colTicker := calledNumbers[board[col][row]]

				if rowMarked && rowTicker == 0 {
					rowMarked = false
				} else if rowMarked && rowTicker > localRowTicker {
					localRowTicker = rowTicker
					localRowNum = board[row][col]
				}

				//fmt.Printf("board: %d, row: %d, col: %d, num: %d, ticker: %d, rowMarked: %t, localRowTicker: %d, localRowNum: %d\n",
				//	boardIdx, row, col, board[row][col], rowTicker, rowMarked, localRowTicker, localRowNum)

				if colMarked && colTicker == 0 {
					colMarked = false
				} else if colMarked && colTicker > localColTicker {
					localColTicker = colTicker
					localColNum = board[col][row]
				}

				//fmt.Printf("board: %d, row: %d, col: %d, num: %d, ticker: %d, colMarked: %t, localColTicker: %d, localColNum: %d\n",
				//	boardIdx, row, col, board[col][row], colTicker, colMarked, localColTicker, localColNum)

			}

			//fmt.Printf("board: %d, winnerTicker: %d, winnerNum: %d, winnerBoardIdx: %d\n",
			//	boardIdx, winnerTicker, winnerNum, winnerBoardIdx)

			if rowMarked && localRowTicker < winnerTicker {
				winnerTicker = localRowTicker
				winnerNum = localRowNum
				winnerBoardIdx = boardIdx
			}

			//fmt.Printf("board: %d, winnerTicker: %d, winnerNum: %d, winnerBoardIdx: %d\n",
			//	boardIdx, winnerTicker, winnerNum, winnerBoardIdx)

			if colMarked && localColTicker < winnerTicker {
				winnerTicker = localColTicker
				winnerNum = localColNum
				winnerBoardIdx = boardIdx
			}

			//fmt.Printf("board: %d, winnerTicker: %d, winnerNum: %d, winnerBoardIdx: %d\n",
			//	boardIdx, winnerTicker, winnerNum, winnerBoardIdx)
		}
	}

	sumUnmarkedNum := 0

	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			num := bingoBoards[winnerBoardIdx][row][col]
			ticker := calledNumbers[num]

			if ticker == 0 {
				sumUnmarkedNum += num
				//fmt.Printf("row: %d, col: %d, num: %d, ticker: %d, sum: %d\n", row, col, num, ticker, sumUnmarkedNum)
			} else if ticker > winnerTicker {
				sumUnmarkedNum += num
				//fmt.Printf("row: %d, col: %d, num: %d, ticker: %d, sum: %d\n", row, col, num, ticker, sumUnmarkedNum)
			}
		}
	}

	fmt.Printf("%d\n", sumUnmarkedNum*winnerNum)

}

func two() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	calledNumbers := make([]int, 120)
	calledNumbersScanned := false
	bingoBoards := make([][][]int, 0)
	bingoBoard := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if !calledNumbersScanned {
			calledNumbersScanned = true
			for idx, strNum := range strings.Split(line, ",") {
				num, _ := strconv.Atoi(strNum)
				calledNumbers[num] = idx + 1
			}
		}

		if strings.Trim(line, " ") == "" {
			bingoBoard = make([][]int, 0)
		} else {
			row := make([]int, 0)
			for _, strNum := range strings.Split(line, " ") {
				num, err := strconv.Atoi(strNum)
				if err != nil {
					continue
				}
				row = append(row, num)
			}
			bingoBoard = append(bingoBoard, row)
			if len(bingoBoard) == 5 {
				bingoBoards = append(bingoBoards, bingoBoard)
			}
		}
	}

	winnerTicker := -1
	winnerNum := -1
	winnerBoardIdx := -1

	for boardIdx, board := range bingoBoards {
		boardTicker := 110
		boardNum := -1

		for row := 0; row < 5; row++ {
			rowMarked := true
			colMarked := true
			localRowNum := -1
			localRowTicker := -1
			localColNum := -1
			localColTicker := -1
			for col := 0; col < 5; col++ {
				rowTicker := calledNumbers[board[row][col]]
				colTicker := calledNumbers[board[col][row]]

				if rowMarked && rowTicker == 0 {
					rowMarked = false
				} else if rowMarked && rowTicker > localRowTicker {
					localRowTicker = rowTicker
					localRowNum = board[row][col]
				}

				//fmt.Printf("board: %d, row: %d, col: %d, num: %d, ticker: %d, rowMarked: %t, localRowTicker: %d, localRowNum: %d\n",
				//	boardIdx, row, col, board[row][col], rowTicker, rowMarked, localRowTicker, localRowNum)

				if colMarked && colTicker == 0 {
					colMarked = false
				} else if colMarked && colTicker > localColTicker {
					localColTicker = colTicker
					localColNum = board[col][row]
				}

				//fmt.Printf("board: %d, row: %d, col: %d, num: %d, ticker: %d, colMarked: %t, localColTicker: %d, localColNum: %d\n",
				//	boardIdx, row, col, board[col][row], colTicker, colMarked, localColTicker, localColNum)

			}

			//fmt.Printf("board: %d, boardTicker: %d, boardNum: %d\n",
			//	boardIdx, boardTicker, boardNum)

			if rowMarked && localRowTicker < boardTicker {
				boardTicker = localRowTicker
				boardNum = localRowNum
			}

			//fmt.Printf("board: %d, boardTicker: %d, boardNum: %d\n",
			//	boardIdx, boardTicker, boardNum)

			if colMarked && localColTicker < boardTicker {
				boardTicker = localColTicker
				boardNum = localColNum
			}

			//fmt.Printf("board: %d, boardTicker: %d, boardNum: %d\n",
			//	boardIdx, boardTicker, boardNum)
		}

		if boardTicker > winnerTicker {
			winnerTicker = boardTicker
			winnerNum = boardNum
			winnerBoardIdx = boardIdx
		}
	}

	sumUnmarkedNum := 0

	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			num := bingoBoards[winnerBoardIdx][row][col]
			ticker := calledNumbers[num]

			if ticker == 0 {
				sumUnmarkedNum += num
				//fmt.Printf("row: %d, col: %d, num: %d, ticker: %d, sum: %d\n", row, col, num, ticker, sumUnmarkedNum)
			} else if ticker > winnerTicker {
				sumUnmarkedNum += num
				//fmt.Printf("row: %d, col: %d, num: %d, ticker: %d, sum: %d\n", row, col, num, ticker, sumUnmarkedNum)
			}
		}
	}

	fmt.Printf("%d\n", sumUnmarkedNum*winnerNum)

}
