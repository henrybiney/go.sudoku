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
		err = nil
		return nil, err
	}

	//check negative values, and values not between 0 and 9
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

func (g *Grid) Row(rowNum int) []int {
	return g.nums[rowNum]
}

//one based
func (g *Grid) Column(colNum int) (col []int, err error) {

	if colNum <= 0 || colNum > 9 {
		err = fmt.Errorf("grid: column number must be between 0 and 9 but got %d", colNum)
		return col, err
	}

	//sort keys
	for _, key := range g.sorted_keys {
		col = append(col, g.nums[key][colNum-1])
	}

	return col, err
}

func (g *Grid) Read(rowNum, colNum int) (val int, err error) {
	if rowNum <= 0 || rowNum > 9 {
		return val, fmt.Errorf("grid: rownum %d must be gt 10 or lt 0 ", rowNum)
	}
	if colNum <= 0 || colNum > 9 {
		return val, fmt.Errorf("grid: colnum %d must be gt 10 or lt 0 ", colNum)
	}

	val = g.nums[rowNum][colNum-1]

	return val, err
}
