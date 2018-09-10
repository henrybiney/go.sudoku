package solver

import (
	"fmt"

	. "go.sudoku/model"
)

const (
	COMPLETE             = 0
	REQUIRES_SPECULATION = 1
	INCONSISTENT         = 2
)

type Solver struct {
	//the current grid that is being solved
	grid Grid
}

func New(g Grid) Solver {
	return Solver{grid: g}
}

func (s *Solver) Solve() Grid {
	grid, state := s.BasicSolve()
	if state == COMPLETE {
		return grid
	}
	return grid
}

//TODO: Add speculative solve; not complete!
func speculativeSolve(g Grid) (grid Grid, state int) {
	solver := New(g)
	grid, state = solver.BasicSolve()
	//if this yields an inconsistent state then
	//return unsolved
	if state == INCONSISTENT {
		fmt.Printf("Inconsistent grid \n")
		grid.PrintGrid()
		//TODO: Add speculation
		return
	}
	if state == COMPLETE {
		fmt.Printf("Solution found \n")
		return
	}
	return
}

//BasicSolve A simple solver. Solves sudoku's which require no speculation
func (s *Solver) BasicSolve() (solution Grid, state int) {
	changing := true
	for changing {

		changing = false
		s.ComputeAllPossibleValues()
		//after we compute all possible values for cells
		//we find the cell with just one possible value
		cellsWithSingleConstraints := s.findCellsWithSingleConstraints()
		if cellsWithSingleConstraints != nil {
			changing = true

			for _, cell := range cellsWithSingleConstraints {
				cellValue, _ := cell.NextPossibleValue()
				fmt.Printf("Updating row %d, col %d with value %d\n", cell.CellRow(), cell.CellColumn(), cellValue)
				cell.NewValue(cellValue)

				s.grid.PrintGrid()
				cell.ResetIterator()

			}
		}
	}

	//cells with no possible constraint; inconsistent state  empty constraints
	//constraint set is empty; this implies that we have solved the sudoku
	s.ComputeAllPossibleValues()
	solution = s.grid
	state = s.checkGridState()

	return solution, state
}

//ComputeAllPossibleValues Computes all constraints in cells
func (s *Solver) ComputeAllPossibleValues() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			//	s.ComputePossibleValuesAt(i, j)
			s.grid.ComputePossibleValuesAt(i, j)
		}

	}
}

//take a look at zero marked cells and
//compute its possible values

//TODO:FIX THIS
func (s Solver) checkGridState() (state int) {
	//check state
	var requiresSpeculation = s.grid.HasRemainingConstraints()

	if requiresSpeculation {
		fmt.Printf("Requires speculation \n")
		state = REQUIRES_SPECULATION
	} else {
		fmt.Printf("Solution found with no speculation \n")
		state = COMPLETE
	}
	return state
}

func (s Solver) findCellsWithSingleConstraints() (cellPositions []*Cell) {
	var grid Grid = s.grid

	for rowNum := 1; rowNum <= 9; rowNum++ {
		for colNum := 1; colNum <= 9; colNum++ {
			cell, _ := grid.ReadCell(rowNum, colNum)
			if len(cell.PossibleValues()) == 1 && cell.CellValue() == 0 {
				cellPositions = append(cellPositions, cell)
			}
		}

	}
	return cellPositions
}
