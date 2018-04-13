package solver

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"strconv"
	"strings"
	. "sudoku/model"
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

	if state == "SPECULATION_REQUIRED" {
		gridStack := stack.New()
		gridStack.Push(grid)
		fmt.Printf("%v, len %d\n", gridStack, gridStack.Len())

		for gridStack.Len() != 0 {
			//set row1,col1 to new value
			//push that grid onto the stack
			//solve it
			var g1 interface{} = gridStack.Pop()
			grid_ := g1.(Grid)
			solver = New(grid_)
			solver.BasicSolve()

			currentConstraints := solver.constraints
			for position, values := range currentConstraints {

				rowColumn := strings.Split(position, ",")
				row, _ := strconv.Atoi(rowColumn[0])
				col, _ := strconv.Atoi(rowColumn[1])

				for possibleValue := range values {

					solver.grid.UpdateValueAt(row, col, possibleValue)
					gridStack.Push(grid_)
					g, s := solver.BasicSolve()

					if s == "INCONSISTENT" {
						fmt.Printf("Inconsistent grid\n")
						g.PrintGrid()
						gridStack.Pop()

					}
					if s == "SPECULATION_REQUIRED" {
						fmt.Printf("Using row=%d, col=%d with value %d requires speculation\n", row, col, possibleValue)
						g.PrintGrid()
						gridStack.Push(g)
					}
				}
			}

			//if that grid leaves an inconsistent state
			//then pop it off the top

			//if it leads to speculation, push it onto the stack

		}
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
			s.ComputePossibleValuesAt(i, j)
		}

	}
}

//take a look at zero marked cells and
//compute its possible values
func (s *Solver) ComputePossibleValuesAt(row, col int) {
	//TODO: CHECK IF THIS IS DUPLICATE LOGIC
	var val int
	var err error

	if val, err = s.grid.Read(row, col); err != nil || val != 0 {
		return
	}
	//value at row, col cannot be in
	// rowValues or colValues or boxValues
	rowValues, errRow := s.grid.Row(row)
	colValues, errCol := s.grid.Column(col)
	boxValues := s.grid.GetBoxValuesAt(row, col)

	//TODO: We need box constraints too

	if errRow != nil || errCol != nil {
		return
	}

	key := fmt.Sprintf("%d,%d", row, col)

	s.constraints[key] = map[int]int{1: 1, 2: 2, 3: 3,
		4: 4, 5: 5, 6: 6,
		7: 7, 8: 8, 9: 9}

	setVal := 0

	//delete row values from the set of possible values
	for i := 0; i < len(rowValues); i++ {
		setVal = rowValues[i]
		delete(s.constraints[key], setVal)
	}

	//delete col values from the set of possible values
	for i := 0; i < len(colValues); i++ {
		setVal = colValues[i]
		delete(s.constraints[key], setVal)
	}

	//delete box values from the set of possible constraints
	for i := 0; i < len(boxValues); i++ {
		setVal = boxValues[i]
		delete(s.constraints[key], setVal)

	}

	//what remains is the set of possible values
}

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
