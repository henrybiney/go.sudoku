package solver_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sort"
	"sudoku/examples"
	"sudoku/model"
	"sudoku/solver"
	"testing"
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
	var (
		s solver.Solver
		g *model.Grid
	)

	BeforeEach(func() {
		g, _ = model.NewGrid(examples.EX1)
		s = solver.New(*g)
	})

	Describe("Testing retrieval of possible values at blank location", func() {
		Context("Given a blank location (1,1) on the EX1 board", func() {
			It("Should return possible values", func() {
				s.ComputePossibleValuesAt(1, 1)
				constraints, _ := s.GetConstraintsAt(1, 1)
				sortedConstraints := sortIntMapByKey(constraints)
				Ω(sortedConstraints).Should(Equal([]int{5, 6, 7, 8}))
			})
		})

		Context("Given a blank location (1,3) on the EX1 board", func() {
			It("It should return possible values {4,9} at this location", func() {

				s.ComputePossibleValuesAt(3, 2)
				constraints, _ := s.GetConstraintsAt(3, 2)
				sortedConstraints := sortIntMapByKey(constraints)
				Ω(sortedConstraints).Should(Equal([]int{4, 9}))
			})

		})
	})

	Describe("Testing basic solve", func() {
		Context("Given a grid which requires speculation, basic specSolve should return state = SPECULATION_REQUIRED", func() {

			It("Should return cells with remaining constraints", func() {
				_, state := s.BasicSolve()
				Expect(state).To(Equal("SPECULATION_REQUIRED"))

			})
		})

		Context("Given a grid which requires speculation, a Speculative solve should return a complete state", func() {

			It("Should return complete state", func() {

				model, _ := model.NewGrid(examples.EX1)
				g, state := solver.SpeculativeSolve(*model)
				g.PrintGrid()
				Expect(state).To(Equal("COMPLETE"))

			})
		})
	})

	Describe("Testing basic solve", func() {
		Context("Given a grid which requires *no* speculation", func() {
			no_specGrid, _ := model.NewGrid(examples.EX2)
			mSolver := solver.New(*no_specGrid)

			It("Should return a complete state", func() {

				grid, state := mSolver.BasicSolve()
				fmt.Printf("Solution found - no speculationi\n")
				grid.PrintGrid()
				Expect(state).To(Equal("COMPLETE"))
			})

		})

	})

})
