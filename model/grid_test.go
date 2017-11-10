package model_test

import (
	. "sudoku/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

		badGrid = []int{0, 0, 0, 2, 1, 2, 2, 0}
	})

	Describe("Testing bad initialization", func() {
		Context("Given an bad array, with less than 81 nums",

			func() {
				It("Should return an error", func() {

					_, err := NewGrid(badGrid)
					Expect(err).To(BeNil())
				})
			})
	})

	Describe("Testing initialization", func() {
		Context("Given an array of 81 nums", func() {

			It("Should return a sudoku grid", func() {

				grid, _ := NewGrid(gridNums)

				Expect(grid != nil).To(BeTrue())

				Expect(grid.RowCount()).To(Equal(9))

				Ω(grid.Row(1)).Should(Equal([]int{0, 3, 0, 0, 0, 0, 0, 4, 0}))

				Ω(grid.Row(4)).Should(Equal([]int{0, 0, 3, 0, 0, 0, 8, 0, 0}))

				Ω(grid.Row(9)).Should(Equal([]int{0, 5, 0, 0, 7, 1, 0, 8, 0}))
			})
		})
	})

	Describe("Testing invalid row and column retrieval", func() {
		Context("Given an invalid column number", func() {
			It("Should return an error", func() {
				grid, _ := NewGrid(gridNums)
				_, err1 := grid.Column(10)
				_, err2 := grid.Column(-1)

				Expect(err1 != nil).To(BeTrue())

				Expect(err2 != nil).To(BeTrue())

			})

		})

	})

	Describe("Testing valid row and column retrieval", func() {
		Context("Given a valid column value", func() {
			It("Should return a valid array of column values", func() {
				grid, _ := NewGrid(gridNums)
				val, err := grid.Read(1, 2)
				Ω(grid.Column(3)).Should(Equal([]int{0, 0, 2, 3, 0, 7, 9, 0, 0}))

				Expect(err == nil).To(BeTrue())
				Expect(val).To(Equal(3))
				val, err = grid.Read(8, 8)
				Expect(val).To(Equal(2))
			})
		})
	})

})
