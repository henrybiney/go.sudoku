package model

import (
	"fmt"
	"sort"
)

// Cell struct defines a cell on the Board
type Cell struct {
	possibleValues []int
	row            int
	col            int
	value          int
	tracker        int
}

// NewCell Initialises a new cell
// Note: this does not set the possible values of the cell.
// Possible values of the cell should be initialized sepearately
func NewCell(row, col, value int) Cell {
	return Cell{row: row, col: col, value: value}

}

func (c Cell) CellValue() int {
	return c.value
}

func (c *Cell) NewValue(n int) {
	c.value = n
}

// SetPossibleValues Sets the possible values of a cell
func (c *Cell) SetPossibleValues(possibleVals []int) {
	c.possibleValues = possibleVals
	sort.Ints(possibleVals)
}

// NextPossibleValue Steps through the set of possible values of cell
// Returns the value at the cell and a boolean false to indicate
// whether we have exhausted the set of possible values at the cell
func (c *Cell) NextPossibleValue() (value int, ok bool) {
	if c.tracker == len(c.possibleValues) {
		ok = false
		return
	}

	value = c.possibleValues[c.tracker]
	c.tracker = c.tracker + 1
	ok = true

	return
}

func (c Cell) PrintConstraints() {

	fmt.Printf("Constraints at row %d , col %d are %v. Cell value: %d, size: %d\n",
		c.row, c.col, c.possibleValues, c.CellValue(), len(c.possibleValues))
}

// ResetIterator resets the reading of the next possible values
// to the beginning
func (c *Cell) ResetIterator() {
	c.tracker = 0
}
