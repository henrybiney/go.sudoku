package solver

import (
	"fmt"

	. "go.sudoku/model"
)

//State : The state of an attempt by the Solver
type State int

const (
	//COMPLETE : Solved; all cells yielded a solution
	COMPLETE State = iota

	//REQUIRES_SPECULATION : One or more cells have 2 or more possible values
	REQUIRES_SPECULATION

	//INCONSISTENT : One or more cells cannot be filled with a value
	INCONSISTENT
)

//Solver : Represents a solver
type Solver struct {
	//the current grid that is being solved
	grid Grid
}

//New :  Create a new solver
func New(g Grid) Solver {
	return Solver{grid: g}
}

//Solve : Attempt to solve using speculation or basicSolve
func (s *Solver) Solve() Grid {
	grid, state := s.BasicSolve()
	if state == COMPLETE {
		return grid
	}

	var stack []*Cell

	var multCells []*Cell = grid.ReadCellsWithConstraints()
	var filled = 81 - len(multCells)

	var idx int = 0

	for filled <= 81 && idx < len(multCells) {

		val := multCells[idx]
		stack = append(stack, val)
		// if can fill cell
		if nextVal, okVal := val.NextPossibleValue(); okVal {
			//if can fill grid with nextVal then
			grid.UpdateValueAt(val.CellRow(), val.CellColumn(), nextVal)
			s.ComputeAllPossibleValues()
			filled++
			idx++
			//fmt.Printf("**SpecSolve****  row %d, col %d  with %d\n", val.CellRow(), val.CellColumn(), val.CellValue())
			//grid.PrintGrid()
			//check bounds,
		} else {
			//
			//fmt.Printf("**SpecSolve Backtracking ... row %d, col %d\n", val.CellRow(), val.CellColumn())
			//fmt.Printf("Possible vals %v\n", val.PossibleValues())
			//fmt.P	rintf("Changing row %d, col %d back to 0")
			val = stack[len(stack)-1]
			//fmt.Printf("Reset previous cell. row %d, col %d to 0\n", val.CellRow(), val.CellColumn())
			val.ResetIterator()
			val.NewValue(0)
			s.ComputeAllPossibleValues()
			stack = stack[:len(stack)-1]

			//if stack is empty and I've tried everything; return inconsistent
			filled--
			idx--
			//grid.PrintGrid()
		}

	}

	return grid
}

//TODO: Add speculative solve; not complete!
func speculativeSolve(g Grid) (grid Grid, state State) {
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
func (s *Solver) BasicSolve() (solution Grid, state State) {
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

func (s Solver) checkGridState() (state State) {
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
	grid := s.grid

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
