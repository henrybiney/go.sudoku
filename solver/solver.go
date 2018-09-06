package solver

import (
	"fmt"
	"strconv"
	"strings"

	. "go.sudoku/model"
)

type Solver struct {

	//we will store grid constraints in a map for fast retrieval
	//key = string  eg.  "1,2" - constraints a row 1, column 2

	//constraint values will also be stored in map for fast updates
	// eg. {9:9, 8:8}
	// key: 1,2 ,  val: {9:9, 8:8}. possible values at position 1,2 are 9 and 8
	constraints map[string]map[int]int

	//the current grid that is being solved
	grid Grid
}

func New(g Grid) Solver {

	return Solver{constraints: make(map[string]map[int]int), grid: g}
}

func SpeculativeSolve(g Grid) (grid Grid, state string) {
	solver := New(g)
	grid, state = solver.BasicSolve()

	//if this yields an inconsistent state then
	//return unsolved
	if state == "INCONSISTENT" {
		fmt.Printf("Inconsistent grid \n")
		grid.PrintGrid()
		return
	}

	if state == "COMPLETE" {
		fmt.Printf("Solution found \n")
		return
	}

	return
}

func (s *Solver) BasicSolve() (solution Grid, state string) {
	changing := true
	for changing {

		changing = false
		s.ComputeAllPossibleValues()
		//after we compute all possible values for cells
		//we find the cell with just one possible value
		cellsWithSingleConstraints := s.findCellsWithSingleConstraints()

		if cellsWithSingleConstraints != nil {
			changing = true

			for _, key := range cellsWithSingleConstraints {

				rowColumn := strings.Split(key, ",")
				row, _ := strconv.Atoi(rowColumn[0])
				col, _ := strconv.Atoi(rowColumn[1])

				for _, v := range s.constraints[key] {
					s.grid.UpdateValueAt(row, col, v)
					delete(s.constraints, key)
				}
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

func (s *Solver) ComputeAllPossibleValues() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			//	s.ComputePossibleValuesAt(i, j)
		}

	}
}

//take a look at zero marked cells and
//compute its possible values

//TODO:FIX THIS
func (s Solver) checkGridState() (state string) {
	//check state
	if len(s.constraints) == 0 {
		state = "COMPLETE"
		return
	}
	inconsistent := false

	for _, consts := range s.constraints {
		if len(consts) == 0 {
			inconsistent = true
			break
		}
	}

	if inconsistent {
		state = "INCONSISTENT"
	} else {

		state = "SPECULATION_REQUIRED"
	}

	return

}

func (s Solver) findCellsWithSingleConstraints() (cellPositions []string) {
	//TODO: Cache this when we compute; do this in ComputePossibleValueAt(i,j)

	for cellPosition := range s.constraints {
		constraints, ok := s.constraints[cellPosition]
		if ok && len(constraints) == 1 {
			cellPositions = append(cellPositions, cellPosition)
		}
	}
	return
}

func (s Solver) GetConstraintsAt(row, col int) (constraints map[int]int, err error) {

	key := fmt.Sprintf("%d,%d", row, col)

	constraints, ok := s.constraints[key]

	if !ok {
		err = fmt.Errorf("Could not find constraints @ row %d, column %d", row, col)
		return
	}

	return
}

// S: use  a stack to track the current state of the Board

// Each cell has a set of possible values Ci = {xi, .. , xn}
// 					if Ci.value != 0

// pick  Ci[0], push it onto the stack
//  if tried all Ci and inconsistent; pop stack

//  if inconsistent; try C[i+1]

//  if consistent:  pick another Cell
