package handler

func convert1Dto2DArray(sudoku [81]int) [9][9]int {
	sudoku2DArray := [9][9]int{}
	currentIndex := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			sudoku2DArray[i][j] = sudoku[currentIndex]
			currentIndex++
		}
	}

	return sudoku2DArray
}
