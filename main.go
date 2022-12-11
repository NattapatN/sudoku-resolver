package main

import (
	"fmt"

	"github.com/NattapatN/sudoku-resolver/handler"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	figure.NewFigure("Sudoku - Resolver", "rectangles", true).Print()
	fmt.Println()
	puzzle := handler.Setup("./resources/sudoku.txt")
	handler.PrintSudokuBoard("Puzzle", puzzle)
	result := handler.Solving()
	handler.PrintSudokuBoard("Solution", result)
}
