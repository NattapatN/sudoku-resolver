package handler

var probabilityAnswer = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func findPropability(sudoku2DArray [9][9]int) [9][9][]int {
	propabilityArray := [9][9][]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku2DArray[i][j] == 0 {
				propabilityArray[i][j] = findPropabilityByPosition(i, j, sudoku2DArray)
			}
		}
	}
	return propabilityArray
}

func findPropabilityByPosition(positionX, positionY int, sudoku [9][9]int) []int {
	propability := []int{}
	//find prop in horizontal
	horizotalArray := formatHorizontalArray(positionX, sudoku)
	propabilityHorizontal := []int{}
	for _, v := range probabilityAnswer {
		if !find(v, horizotalArray) {
			propabilityHorizontal = append(propabilityHorizontal, v)
		}
	}
	//find prop in vertical
	verticalArray := formatVerticalArray(positionY, sudoku)
	propabilityVertical := []int{}
	for _, v := range probabilityAnswer {
		if !find(v, verticalArray) {
			propabilityVertical = append(propabilityVertical, v)
		}
	}
	//find prop in block (3x3)
	blockArray := formatBlockArray(positionX, positionY, sudoku)
	propabilityBlock := []int{}
	for _, v := range probabilityAnswer {
		if !find(v, blockArray) {
			propabilityBlock = append(propabilityBlock, v)
		}
	}
	//find intersection
	propability = findIntersection(propabilityHorizontal, propabilityVertical)
	propability = findIntersection(propability, propabilityBlock)

	return propability
}

func find(number int, dataArray [9]int) bool {
	for _, v := range dataArray {
		if v == number {
			return true
		}
	}
	return false
}

func findIntersection(a, b []int) []int {
	result := []int{}
	for _, v := range a {
		if findInSlice(v, b) {
			result = append(result, v)
		}
	}

	return result
}

func findInSlice(number int, dataArray []int) bool {
	for _, v := range dataArray {
		if v == number {
			return true
		}
	}
	return false
}
