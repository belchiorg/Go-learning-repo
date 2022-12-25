package main

func BreakChocolate(n, m int) int {
	if n == 1 && m == 1 {
		return 0
	}
	if n < 1 || m < 1 {
		return 0
	}

	var verticalSlices = n - 1

	var horizSlices = m - 1

	return n*horizSlices + verticalSlices
}
