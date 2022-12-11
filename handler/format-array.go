package handler

func formatHorizontalArray(positionX int, sudoku [9][9]int) [9]int {
	horizontalArray := [9]int{}
	for i := 0; i < 9; i++ {
		horizontalArray[i] = sudoku[positionX][i]
	}
	return horizontalArray
}

func formatVerticalArray(positionY int, sudoku [9][9]int) [9]int {
	verticalArray := [9]int{}
	for i := 0; i < 9; i++ {
		verticalArray[i] = sudoku[i][positionY]
	}
	return verticalArray
}

func formatBlockArray(positionX, positionY int, sudoku [9][9]int) [9]int {
	blockArray := [9]int{}
	currentIndex := 0
	for i := 3 * (positionX / 3); i < (3*(positionX/3) + 3); i++ {
		for j := 3 * (positionY / 3); j < (3*(positionY/3) + 3); j++ {
			blockArray[currentIndex] = sudoku[i][j]
			currentIndex++
		}
	}
	return blockArray
}
