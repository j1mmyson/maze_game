package stage_test

import (
	"testing"
)

func TestLev1Map(t *testing.T) {
	length := 20
	m := make([][]byte, length)
	for i := 0; i < length; i++ {
		m[i] = make([]byte, length)
	}

	for rowInd, row := range m {
		for colInd := range row {
			if rowInd == 0 || rowInd == 19 || colInd == 0 || colInd == 19 {
				m[rowInd][colInd] = '#'
			}
		}
	}
}
