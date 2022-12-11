package handler

import "fmt"

func PrintSudokuBoard(text string, result [9][9]int) {
	fmt.Println(text, ":")
	for i := 0; i < 9; i++ {
		if i/3 != 0 && i%3 == 0 {
			fmt.Println("- - - + - - - + - - -")
		}
		for j := 0; j < 9; j++ {
			if j/3 != 0 && j%3 == 0 {
				fmt.Print("| ")
			}
			if result[i][j] == 0 {
				fmt.Print(". ")
			} else {
				fmt.Print(result[i][j], " ")
			}
		}
		fmt.Println("")
	}
}
