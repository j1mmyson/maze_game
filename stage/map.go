package stage

import "math/rand"

// Lev1Map returns map of level1.
func Lev1Map() ([][]byte, int) {
	r := rand.New(rand.NewSource(25))
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

	for i := 0; i < 100; i++ {
		m[r.Intn(length-1)+1][r.Intn(length-1)+1] = '#'
	}

	m[length-2][length-2] = '@'

	// display.Map(m)
	return m, length
}
