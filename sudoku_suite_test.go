package sudoku_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSudoku(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sudoku Suite")
}
