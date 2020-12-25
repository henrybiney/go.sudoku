package solver

import (
	"fmt"
	"sort"
	"testing"

	"go.sudoku/examples"
	"go.sudoku/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSolver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Solver Suite")
}

func sortIntMapByKey(constraints map[int]int) (sortedKeys []int) {

	for key := range constraints {
		sortedKeys = append(sortedKeys, key)

	}
	sort.Ints(sortedKeys)
	return
}

var _ = Describe("Sudoku Solver Suite", func() {

	Describe("Basic Solve Test: ", func() {
		Context("Given a grid, Metro @ 12/09/18 which requires *NO* speculation", func() {
			no_specGrid, _ := model.NewGrid(examples.METRO_12_09_18_EASY)
			mSolver := NewSolver(*no_specGrid)

			It("Should return a complete state", func() {

				grid, state := mSolver.BasicSolve()
				grid.PrintGrid()
				Expect(state).To(Equal(COMPLETE))
			})

			Context("Given a grid, Metro @ 21/03/18 which requires *NO* speculation", func() {
				no_specGrid, _ := model.NewGrid(examples.METRO_21_03_18_EASY)
				mSolver := NewSolver(*no_specGrid)

				It("Should return a complete state", func() {
					grid := mSolver.Solve()
					fmt.Println("*************************")
					grid.PrintGrid()

				})
			})

		})
		Context("Given a grid which requires speculation:", func() {
			speculativeGrid, _ := model.NewGrid(examples.EVIL_SUDOKU)
			mSolver := NewSolver(*speculativeGrid)

			It("Should return a 'REQUIRES_SPECULATION' state - EvilSudoku", func() {
				fmt.Printf("EvilSudoku - https://www.websudoku.com/?level=4&set_id=5866622639 \n")
				grid := mSolver.Solve()
				grid.PrintGrid()
				Expect(grid).ToNot(BeNil())
			})
		})
		Context("Given a grid :", func() {
			speculativeGrid, _ := model.NewGrid(examples.METRO_17_09_18_EASY)
			mSolver := NewSolver(*speculativeGrid)

			It("Should return a complete state", func() {
				fmt.Printf("Metro , 17/09/2018 Easy  \n")
				grid := mSolver.Solve()
				grid.PrintGrid()
				//Expect(state).To(Equal(solver.REQUIRES_SPECULATION))
			})
		})

		Context("Given a grid :", func() {
			speculativeGrid, _ := model.NewGrid(examples.METRO_17_09_18_MODERATE)
			mSolver := NewSolver(*speculativeGrid)

			It("Should return a complete state", func() {
				fmt.Printf("Metro , 17/09/2018 MODERATE  \n")
				grid := mSolver.Solve()
				grid.PrintGrid()
				//Expect(state).To(Equal(solver.REQUIRES_SPECULATION))
			})
		})

		Context("Given a grid :", func() {
			speculativeGrid, _ := model.NewGrid(examples.METRO_17_09_18_CHALLENGING)
			mSolver := NewSolver(*speculativeGrid)

			It("Should return a complete state", func() {
				fmt.Printf("Metro , 17/09/2018 CHALLENGING  \n")
				grid := mSolver.Solve()
				grid.PrintGrid()
				//Expect(state).To(Equal(solver.REQUIRES_SPECULATION))
			})
		})

		Context("Given a grid which requires speculation:", func() {
			speculativeGrid, _ := model.NewGrid(examples.EVIL_SUDOKU_2)
			mSolver := NewSolver(*speculativeGrid)

			It("Should return a 'REQUIRES_SPECULATION' state - EvilSudoku 2 ", func() {
				fmt.Printf("EvilSudoku 2 - https://www.websudoku.com/?level=4&set_id=5866622639 \n")
				grid := mSolver.Solve()
				grid.PrintGrid()
				Expect(grid).ToNot(BeNil())
				//Expect(state).To(Equal(solver.REQUIRES_SPECULATION))
			})
		})
	})

	// Describe("Testing retrieval of possible values at blank location", func() {
	// 	Context("Given a blank location (1,1) on the EX1 board", func() {
	// 		It("Should return possible values", func() {
	// 			//	s.ComputePossibleValuesAt(1, 1)
	// 			constraints, _ := s.GetConstraintsAt(1, 1)
	// 			sortedConstraints := sortIntMapByKey(constraints)
	// 			Ω(sortedConstraints).Should(Equal([]int{5, 6, 7, 8}))
	// 		})
	// 	})

	// 	Context("Given a blank location (1,3) on the EX1 board", func() {
	// 		It("It should return possible values {4,9} at this location", func() {

	// 			//s.ComputePossibleValuesAt(3, 2)
	// 			constraints, _ := s.GetConstraintsAt(3, 2)
	// 			sortedConstraints := sortIntMapByKey(constraints)
	// 			Ω(sortedConstraints).Should(Equal([]int{4, 9}))

	// 		})

	// 	})
	// })

	// Describe("Testing basic solve", func() {
	// 	Context("Given a grid which requires speculation, basic specSolve should return state = SPECULATION_REQUIRED", func() {

	// 		It("Should return cells with remaining constraints", func() {
	// 			_, state := s.BasicSolve()
	// 			Expect(state).To(Equal("SPECULATION_REQUIRED"))

	// 		})
	// 	})

	// Context("Given a grid which requires speculation, a Speculative solve should return a complete state", func() {

	// 	It("Should return complete state", func() {

	// 		model, _ := model.NewGrid(examples.EX2)
	// 		//g, state := solver.SpeculativeSolve(*model)
	// 		//g.PrintGrid()
	// 		//Expect(state).To(Equal("COMPLETE"))

	// 		solver.BasicSolve(*model)

	// 	})
	// })

})
