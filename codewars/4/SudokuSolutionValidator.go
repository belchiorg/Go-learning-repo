package main

func ValidateSolution(m [][]int) bool {
	return checkRows(m) && checkCols(m) && checkSquares(m)
}

func checkRows(m [][]int) bool {
	for _, row := range m {
		for i, n := range row {
			for j := i + 1; j < 9; j++ {
				if n == row[j] {
					return false
				}
			}
		}
	}
	return true
}

func checkCols(m [][]int) bool {
	for col := 0; col < 9; col++ {
		for row := 0; row < 9; row++ {
			if m[row][col] == 0 {
				continue
			}
			for row_aux := row + 1; row_aux < 9; row_aux++ {
				if m[row_aux][col] == m[row][col] {
					return false
				}
			}
		}
	}
	return true
}

func checkSquares(m [][]int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !checkSquare(m, i, j) {
				return false
			}
		}
	}
	return true
}

func checkSquare(m [][]int, i int, j int) bool {
	arr := []int{}

	for squareRow := 0; squareRow < 3; squareRow++ {
		for squareCol := 0; squareCol < 3; squareCol++ {
			arr = append(arr, m[squareRow+(i*3)][squareCol+(j*3)])
		}
	}

	for i, n := range arr {
		for j := i + 1; j < 9; j++ {
			if n == arr[j] {
				return false
			}
		}
	}

	return true

}
