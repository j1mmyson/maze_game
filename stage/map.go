package stage

import "math/rand"

type location struct {
	x int
	y int
}

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
			if rowInd == 0 || rowInd == length-1 || colInd == 0 || colInd == length-1 {
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

// BossMap returns map of Boss stage.
func BossMap() ([][]byte, int) {
	r := rand.New(rand.NewSource(33))
	length := 40
	m := make([][]byte, length)
	monsters := make([]location, 3)
	keys := make([]location, 3)

	for i := 0; i < length; i++ {
		m[i] = make([]byte, length)
	}

	for rowInd, row := range m {
		for colInd := range row {
			if rowInd == 0 || rowInd == length-1 || colInd == 0 || colInd == length-1 {
				m[rowInd][colInd] = '#'
			}
		}
	}

	for i := 0; i < 3; i++ {
		monsters[i].x = r.Intn(36) + 5
		monsters[i].y = r.Intn(38)
		if m[monsters[i].x][monsters[i].y] == 0 {
			m[monsters[i].x][monsters[i].y] = 'M'
		} else {
			i = i - 1
		}
	}

	for i := 0; i < 3; i++ {
		keys[i].x = r.Intn(36) + 2
		keys[i].y = r.Intn(38)
		if m[keys[i].x][keys[i].y] == 0 {
			m[keys[i].x][keys[i].y] = 'k'
		} else {
			i = i - 1
		}
	}

	for i := 0; i < 300; i++ {
		m[r.Intn(length-1)+1][r.Intn(length-1)+1] = '#'
	}

	return m, length
}
