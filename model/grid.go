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
		return
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

	g = &Grid{nums: rep, sorted_keys: keys}
	return
}

func (g *Grid) PrintGrid() {
	fmt.Printf("-------------------------------------")
	fmt.Printf("\n")
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

func (g *Grid) RowCount() int {
	return len(g.nums)
}

func (g *Grid) Row(rowNum int) (row []int, err error) {
	if rowNum <= 0 || rowNum > 9 {
		return nil, fmt.Errorf("The row at %d does not exist", rowNum)
	}
	row = g.nums[rowNum]
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
		err = fmt.Errorf("There was an error reading constraints at row %d, col %d",
			rowNum, colNum)
		return constraints, err

	}

	return
}

func (g *Grid) GetBoxValuesAt(rowNum, colNum int) (boxValues []int) {

	bounds := getBox(rowNum, colNum)
	var val int = 0
	for i := bounds.rowSt; i <= bounds.rowEnd; i++ {
		for j := bounds.colSt; j <= bounds.colEnd; j++ {
			if val, _ = g.Read(i, j); val != 0 {

				boxValues = append(boxValues, val)
			}
		}
	}

	return
}
