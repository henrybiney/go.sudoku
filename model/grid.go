package model

import (
	"fmt"
	"sort"
)

// Grid represents a sudoku grid
type Grid struct {
	nums       map[int][]Cell
	sortedKeys []int
}

// NewGrid : initialises a sudoku grid
// Returns a pointer to the initialized grid if values on the grid are ok
// together with the status of the err
// err is Nil if initialization was ok or non-nil if otherwise
func NewGrid(vals []int) (g *Grid, err error) {

	if vals == nil || len(vals) != 81 {
		err = fmt.Errorf("Board initialization error. Board is nil or values less than 81")
		return
	}

	rep := make(map[int][]Cell)

	//read row based array
	//row 1 starts array index 0, row 2 at index 11
	for i, row, col := 0, 1, 1; i < len(vals); i++ {
		rep[row] = append(rep[row], NewCell(row, col, vals[i]))
		col++

		if (i+1)%9 == 0 {
			row++
			col = 1
		}
	}

	var keys []int
	for k := range rep {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	g = &Grid{nums: rep, sortedKeys: keys}

	g.computeAllConstraints()
	//g.nums[rowNum][colNum-1].SetPossibleValues(possibleVals)
	return
}

func (g *Grid) computeAllConstraints() {
	for row := 1; row <= 9; row++ {
		for col := 1; col <= 9; col++ {
			g.ComputePossibleValuesAt(row, col)
		}
	}
}

// ComputePossibleValuesAt  : returns some value at the row
func (g *Grid) ComputePossibleValuesAt(row, col int) {

	if val, err := g.Read(row, col); err != nil || val != 0 {
		return
	}

	rowValues, errRow := g.Row(row)
	colValues, errCol := g.Column(col)
	boxValues := g.GetBoxValuesAt(row, col)

	if errRow != nil || errCol != nil {
		return
	}

	possibleValues := map[int]int{1: 1, 2: 2, 3: 3,
		4: 4, 5: 5, 6: 6,
		7: 7, 8: 8, 9: 9}
	//delete row values from the set of possible values
	for i := 0; i < len(rowValues); i++ {
		delete(possibleValues, rowValues[i])
	}
	//delete col values from the set of possible values
	for i := 0; i < len(colValues); i++ {
		delete(possibleValues, colValues[i])
	}
	//delete box values from the set of possible constraints
	for i := 0; i < len(boxValues); i++ {
		delete(possibleValues, boxValues[i])
	}
	var possValArr []int

	for i := range possibleValues {
		possValArr = append(possValArr, i)
	}

	g.nums[row][col-1].SetPossibleValues(possValArr)

}

func (g *Grid) Read(rowNum, colNum int) (val int, err error) {

	if !isValidBound(rowNum) || !isValidBound(colNum) {
		return val, fmt.Errorf("grid: rownum %d or col %d must be gt 10 or lt 0 ", rowNum, colNum)
	}

	val = g.nums[rowNum][colNum-1].CellValue()

	return
}

// Row Returns a Row given the row number
// An error is returned if the row num less than or equal 0
// or greater than 9
func (g *Grid) Row(rowNum int) (row []int, err error) {
	if !isValidBound(rowNum) {
		return nil, fmt.Errorf("The row at %d does not exist", rowNum)
	}
	cells := g.nums[rowNum]
	for _, cell := range cells {
		row = append(row, cell.CellValue())
	}
	return
}

// Column Returns a Column given the row number
// An error is returned if the row num less than or equal 0
// or greater than 9
func (g *Grid) Column(colNum int) (col []int, err error) {

	if !isValidBound(colNum) {
		err = fmt.Errorf("grid: column number must be between 0 and 9 but got %d", colNum)
		return
	}

	//sort keys
	for _, key := range g.sortedKeys {
		col = append(col, g.nums[key][colNum-1].CellValue())
	}

	return
}

//GetBoxValuesAt Returns bounding values in a sudoku box (3x3 box)
func (g *Grid) GetBoxValuesAt(rowNum, colNum int) (boxValues []int) {

	bounds := getBox(rowNum, colNum)
	var val int
	for i := bounds.rowSt; i <= bounds.rowEnd; i++ {
		for j := bounds.colSt; j <= bounds.colEnd; j++ {

			if val, _ = g.Read(i, j); val != 0 {
				boxValues = append(boxValues, val)
			}
		}
	}

	return
}

// ShowConstraints : print cells and their possible values
func (g Grid) ShowConstraints() {
	for _, row := range g.sortedKeys {
		for _, cell := range g.nums[row] {
			cell.PrintConstraints()

		}
	}
}

// PrintGrid :prints the sudoku grid
func (g *Grid) PrintGrid() {
	fmt.Printf("-------------------------------------\n")
	for i := 1; i <= 9; i++ {

		for j := 1; j <= 9; j++ {
			v, _ := g.Read(i, j)
			fmt.Printf(" %d |", v)
		}
		if i%3 == 0 {
			fmt.Print("\n")
			fmt.Printf("------------------------------------")

		}

		fmt.Printf("\n")

	}
	fmt.Printf("\n")

}

// RowCount : returns the number of rows in the grid
func (g *Grid) RowCount() int {
	return len(g.nums)
}

//ReadCell : reads a  cell on a grid given the rowNum and column number
func (g Grid) ReadCell(row, col int) (cell *Cell, err error) {
	if !isValidBound(row) || !isValidBound(col) {
		return cell, fmt.Errorf("grid: rownum %d or col %d must be gt 10 or lt 0 ", row, col)
	}
	cell = &g.nums[row][col-1]
	return
}

func isValidBound(position int) bool {
	return !(position <= 0 || position > 9)

}

// UpdateValueAt updates a value at row and column
func (g *Grid) UpdateValueAt(rowNum, colNum, newValue int) (err error) {
	if !isValidBound(rowNum) || !isValidBound(colNum) {
		err = fmt.Errorf("Invalid row %d or column %d on grid", rowNum, colNum)
		return
	}
	g.nums[rowNum][colNum-1].NewValue(newValue)
	return
}

//ConstraintsAt  : Is this supposed to be in the solver module or here?
func (g *Grid) ConstraintsAt(rowNum, colNum int) (constraints []int, err error) {

	if _, err = g.Read(rowNum, colNum); err != nil {
		err = fmt.Errorf("There was an error reading constraints at row %d, col %d",
			rowNum, colNum)

		return constraints, err

	}
	constraints, err = g.nums[rowNum][colNum-1].PossibleValues(), nil

	return
}

func (g *Grid) ReadCellsWithConstraints() (cells []*Cell) {

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			cell, _ := g.ReadCell(i, j)
			if cell.CellValue() == 0 && len(cell.PossibleValues()) > 1 {
				cells = append(cells, cell)
			}
		}
	}
	return
}

//HasRemainingConstraints : Check if a cell has more constraints
func (g *Grid) HasRemainingConstraints() bool {

	var hasRemainingConstraints = false
	for rowNum := 1; rowNum <= 9; rowNum++ {
		for colNum := 1; colNum <= 9; colNum++ {
			cell, _ := g.ReadCell(rowNum, colNum)
			val := cell.CellValue()
			hasRemainingConstraints = hasRemainingConstraints || val == 0
		}
	}
	return hasRemainingConstraints
}

//CopyGrid : Make a deep copy of the Grid
//This looks rather ugly;
func (g Grid) CopyGrid() (newGrid Grid) {

	newGrid.sortedKeys = make([]int, len(g.sortedKeys))
	newGrid.nums = make(map[int][]Cell)
	copy(newGrid.sortedKeys, g.sortedKeys)

	for k, v := range g.nums {
		newGrid.nums[k] = make([]Cell, len(v))
		copy(newGrid.nums[k], v)
	}
	return

}
