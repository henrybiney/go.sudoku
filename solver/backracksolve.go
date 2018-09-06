package solver

import (
	"fmt"

	. "go.sudoku/model"
)

func BasicSolve(grid Grid) (solution Grid) {

	fmt.Printf("Before solving*******\n")
	grid.PrintGrid()

	solvingSingleConstraints := true

	for solvingSingleConstraints {

		solvingSingleConstraints = false
		for i := 1; i <= 9; i++ {
			for j := 1; j <= 9; j++ {
				cell, _ := grid.ReadCell(i, j)
				cellValues := cell.PossibleValues()
				if cell.CellValue() == 0 && len(cellValues) == 1 {
					cell.NewValue(cellValues[0])
					solvingSingleConstraints = true
					grid.PrintGrid()
					fmt.Printf("***\n")
				}

			}

		}

	}
	fmt.Printf("After solving*******\n")
	grid.PrintGrid()
	return

}
