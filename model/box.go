package model

type box struct {
	rowSt  int
	rowEnd int
	colSt  int
	colEnd int
}

func getBox(row, col int) (bounds box) {

	minRow, maxRow := 0, 0

	if row >= 1 && row <= 3 {
		minRow, maxRow = 1, 3

		if col <= 3 {
			bounds = box{minRow, maxRow, 1, 3}
			return
		}
		if col <= 6 {
			bounds = box{minRow, maxRow, 4, 6}
			return
		}
		if col <= 9 {
			bounds = box{minRow, maxRow, 7, 9}
			return
		}

	}

	if row >= 4 && row <= 6 {
		minRow, maxRow = 4, 6
		if col <= 3 {
			bounds = box{minRow, maxRow, 1, 3}
			return
		}
		if col <= 6 {
			bounds = box{minRow, maxRow, 4, 6}
			return
		}
		if col <= 9 {
			bounds = box{minRow, maxRow, 7, 9}
			return
		}
	}

	if row >= 7 && row <= 9 {
		minRow, maxRow = 7, 9
		if col <= 3 {
			bounds = box{minRow, maxRow, 1, 3}
			return
		}
		if col <= 6 {
			bounds = box{minRow, maxRow, 4, 6}
			return
		}
		if col <= 9 {
			bounds = box{minRow, maxRow, 7, 9}
			return
		}

	}
	return
}
