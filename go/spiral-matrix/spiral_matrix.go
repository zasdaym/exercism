package spiralmatrix

func SpiralMatrix(size int) [][]int {
	result := make([][]int, size)
	for i := range result {
		result[i] = make([]int, size)
	}

	var (
		row int
		col int
		dir int
	)
	for i := 1; i <= size*size; i++ {
		result[row][col] = i
		switch dir {
		case 0:
			if col+1 == size || result[row][col+1] != 0 {
				dir = 1
				row++
			} else {
				col++
			}
		case 1:
			if row+1 == size || result[row+1][col] != 0 {
				dir = 2
				col--
			} else {
				row++
			}
		case 2:
			if col-1 < 0 || result[row][col-1] != 0 {
				dir = 3
				row--
			} else {
				col--
			}
		case 3:
			if row-1 < 0 || result[row-1][col] != 0 {
				dir = 0
				col++
			} else {
				row--
			}
		}
	}
	return result
}
