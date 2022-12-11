package handler

import (
	"os"
	"strconv"
)

func ReadFile(sudoku string) [81]int {
	data, err := os.ReadFile(sudoku)
	if err != nil {
		panic(err)
	}
	return convertToArray(data)
}
func convertToArray(data []byte) [81]int {
	result := [81]int{}
	currentIndex := 0
	for _, v := range data {
		if string(v) == " " {
			result[currentIndex] = 0
		} else if n, err := strconv.Atoi(string(v)); err == nil {
			result[currentIndex] = n
		} else {
			continue
		}
		currentIndex++
	}
	return result
}
