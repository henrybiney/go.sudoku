package model

import (
	"fmt"
	"sort"
)

type Grid struct {
	nums        map[int][]int
	sorted_keys []int
}

func NewGridFromFile(file string) (g *Grid, err error) {

	return nil, nil
}

func NewGrid(vals []int) (g *Grid, err error) {

	if vals == nil || len(vals) != 81 {
		err = fmt.Errorf("Board initialization error. Board is nil or values less than 81")
		return nil, err
	}

	//TODO:check negative values, and values not between 0 and 9
	rep := make(map[int][]int)
	for i, j, row := 0, 0, 1; i < len(vals); i++ {

		if (i+1)%9 == 0 {

			rep[row] = vals[j:(i + 1)]
			row += 1
			j = i + 1
		}
	}

	var keys []int
	for k := range rep {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return &Grid{nums: rep, sorted_keys: keys}, err
}

func (g *Grid) RowCount() int {
	return len(g.nums)
}

func (g *Grid) Row(rowNum int) (row []int, err error) {
	if rowNum <= 0 || rowNum > 9 {
		return nil, fmt.Errorf("The row at %d does not exist", rowNum)
	}
	row, err = g.nums[rowNum], nil
	return
}

//one based
func (g *Grid) Column(colNum int) (col []int, err error) {

	if colNum <= 0 || colNum > 9 {
		err = fmt.Errorf("grid: column number must be between 0 and 9 but got %d", colNum)
		return
	}

	//sort keys
	for _, key := range g.sorted_keys {
		col = append(col, g.nums[key][colNum-1])
	}

	return
}

func (g *Grid) Read(rowNum, colNum int) (val int, err error) {

	if rowNum <= 0 || rowNum > 9 {
		return val, fmt.Errorf("grid: rownum %d must be gt 10 or lt 0 ", rowNum)
	}

	if colNum <= 0 || colNum > 9 {
		return val, fmt.Errorf("grid: colnum %d must be gt 10 or lt 0 ", colNum)
	}

	val = g.nums[rowNum][colNum-1]

	return
}

func (g *Grid) UpdateValueAt(rowNum, colNum, newValue int) (err error) {
	if rowNum <= 0 || rowNum > 9 {
		err = fmt.Errorf("Invalid row %d or column %d on grid", rowNum, colNum)
		return
	}

	g.nums[rowNum][colNum-1] = newValue
	return
}

//TODO: Is this supposed to be in the solver module or here?
func (g *Grid) ConstraintsAt(rowNum, colNum int) (constraints []int, err error) {

	val, err := g.Read(rowNum, colNum)

	if val != 0 || err != nil {
		err = fmt.Errorf("There was an error reading constraints at row %d, col %d", rowNum, colNum)
		return constraints, err

	}

	return nil, err
}
