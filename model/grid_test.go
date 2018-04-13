package model_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sort"
	. "sudoku/model"
)

var _ = Describe("Grid Model tests", func() {

	var (
		gridNums []int
		badGrid  []int
	)
	BeforeEach(func() {
		gridNums = []int{
			0, 3, 0, 0, 0, 0, 0, 4, 0,
			0, 1, 0, 0, 9, 7, 0, 5, 0,
			0, 0, 2, 5, 0, 8, 6, 0, 0,
			0, 0, 3, 0, 0, 0, 8, 0, 0,
			9, 0, 0, 0, 0, 4, 3, 0, 0,
			0, 0, 7, 6, 0, 0, 0, 0, 4,
			0, 0, 9, 8, 0, 5, 4, 0, 0,
			0, 7, 0, 0, 0, 0, 0, 2, 0,
			0, 5, 0, 0, 7, 1, 0, 8, 0}

		badGrid = []int{0, 0, 0, 2, 1, 2, 2, 0, 88}
	})

	Describe("Testing initialization", func() {
		Context("Given an bad array, with less than 81 nums", func() {
			It("Should return an error", func() {

				_, err := NewGrid(badGrid)
				Expect(err).To(Not(BeNil()))
			})
		})

		Context("Given an array of 81 nums", func() {
			It("Should return a sudoku grid", func() {

				grid, _ := NewGrid(gridNums)

				Expect(grid != nil).To(BeTrue())

				Expect(grid.RowCount()).To(Equal(9))

				row1, _ := grid.Row(1)
				row4, _ := grid.Row(4)
				row9, _ := grid.Row(9)

				Ω(row1).Should(Equal([]int{0, 3, 0, 0, 0, 0, 0, 4, 0}))

				Ω(row4).Should(Equal([]int{0, 0, 3, 0, 0, 0, 8, 0, 0}))

				Ω(row9).Should(Equal([]int{0, 5, 0, 0, 7, 1, 0, 8, 0}))
			})
		})
	})

	Describe("Testing valid invalid row and column retrieval", func() {
		Context("Given an invalid row number", func() {
			It("Should return an error", func() {

				grid, _ := NewGrid(gridNums)

				row0, err := grid.Row(0)
				Expect(err != nil).To(BeTrue())
				Expect(row0 == nil).To(BeTrue())

				row10, err := grid.Row(10)
				Expect(row10 == nil).To(BeTrue())
				Expect(err != nil).To(BeTrue())
			})
		})

		Context("Given an invalid column number", func() {
			It("Should return an error", func() {
				grid, _ := NewGrid(gridNums)
				_, err1 := grid.Column(10)
				_, err2 := grid.Column(-1)

				Expect(err1 != nil).To(BeTrue())

				Expect(err2 != nil).To(BeTrue())

			})

		})

		Context("Given a valid column value", func() {
			It("Should return a valid array of column values", func() {
				grid, _ := NewGrid(gridNums)
				val, err := grid.Read(1, 2)
				Ω(grid.Column(3)).Should(Equal([]int{0, 0, 2, 3, 0,
					7, 9, 0, 0}))

				Expect(err == nil).To(BeTrue())
				Expect(val).To(Equal(3))
				val, err = grid.Read(8, 8)
				Expect(val).To(Equal(2))

				val, err = grid.Read(9, 5)
				Expect(val).To(Equal(7))
			})
		})
	})

	Describe("Testing retrieval of constraints at grid position", func() {
		Context("Given a grid position with no constraints", func() {
			It("Should return an error", func() {
				grid, _ := NewGrid(gridNums)
				//the value at row 1, colum 2
				constraints, err := grid.ConstraintsAt(1, 2)
				Expect(err != nil).To(BeTrue())
				Expect(constraints == nil).To(BeTrue())
			})
		})

		Context("Given a grid position not on the board", func() {
			It("Should return an error", func() {

				grid, _ := NewGrid(gridNums)
				_, err := grid.ConstraintsAt(-1, 12)
				Expect(err != nil).To(BeTrue())
			})
		})

	})
	Describe("Testing update of a value on the sudoku grid", func() {
		Context("Given a new value on a valid grid position", func() {
			It("It should update grid position to a new value", func() {
				grid, _ := NewGrid(gridNums)
				grid.UpdateValueAt(1, 5, 7)
				updatedValue, _ := grid.Read(1, 5)
				Expect(updatedValue).To(Equal(7))
			})
		})

		Context("Given a new value at an *invalid* grid position", func() {
			It("Should return error", func() {
				grid, _ := NewGrid(gridNums)
				err := grid.UpdateValueAt(-1, 12, 5)

				Expect(err != nil).To(BeTrue())
			})
		})

	})

	Describe("A test to find a bounding box values at a position", func() {
		Context("Given a valid grid position", func() {
			It("Should return non zero bounding values", func() {
				grid, _ := NewGrid(gridNums)
				vals := grid.GetBoxValuesAt(8, 8)
				sort.Ints(vals)
				Ω(vals).Should(Equal([]int{2, 4, 8}))
			})
		})
	})
})
