package handler

import (
	"fmt"
)

var sudoku2DArray [9][9]int

func Setup(path string) [9][9]int {
	sudokuArray := ReadFile(path)
	sudoku2DArray = convert1Dto2DArray(sudokuArray)
	return sudoku2DArray
}

func Solving() [9][9]int {
	//find probability
	isSolved := false
	currentRound := 0
	for !isSolved && currentRound < 1000 {
		probabilityArray := findPropability(sudoku2DArray)

		beforeFillBoard := sudoku2DArray
		//fill the board
		isSolved = fillBoard(probabilityArray)
		if checkBoardWasSameBefore(beforeFillBoard) {
			pickAndFillBoard(probabilityArray)
		}
		currentRound++
	}
	fmt.Println("-----------------------")
	fmt.Println("total run :", currentRound, " rounds. ")
	fmt.Println("-----------------------")
	return sudoku2DArray
}

func fillBoard(probabilityArray [9][9][]int) bool {
	isSolved := true
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if len(probabilityArray[i][j]) == 1 {
				sudoku2DArray[i][j] = probabilityArray[i][j][0]
			} else if len(probabilityArray[i][j]) > 1 {
				isSolved = false
			}
		}
	}
	return isSolved
}

func checkBoardWasSameBefore(beforeBoard [9][9]int) bool {
	emptyLengthBefore := 0
	emptyLengthAfter := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if beforeBoard[i][j] == 0 {
				emptyLengthBefore++
			}
			if sudoku2DArray[i][j] == 0 {
				emptyLengthAfter++
			}
		}
	}
	return emptyLengthBefore == emptyLengthAfter
}

func pickAndFillBoard(probabilityArray [9][9][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if len(probabilityArray[i][j]) > 1 {
				sudoku2DArray[i][j] = probabilityArray[i][j][0]
				return
			}
		}
	}
}
