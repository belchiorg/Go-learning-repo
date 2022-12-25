package main

import "math"

func ListSquared(m, n int) [][]int {
	var arr [][]int
	for i := m; i <= n; i++ {
		if i == 1 {
			arr = append(arr, []int{1, 1})
			continue
		}
		val := isSquared(i)
		if val != -1 {
			arr = append(arr, []int{i, val})
		}
	}

	if len(arr) == 0 {
		return [][]int{}
	}
	return arr
}

func isSquared(n int) int {
	arr := []int{1, n}
	lim := n / 2
	for i := 2; i < lim; i++ {
		if n%i == 0 {
			arr = append(arr, i)
			if i != n/i {
				arr = append(arr, n/i)
			}
			lim = n / i
		}
	}

	sum := 0
	for _, n := range arr {
		sum += n * n
	}

	if isIntegral(math.Sqrt(float64(sum))) {
		return sum
	} else {
		return -1
	}

}

func isIntegral(val float64) bool {
	return val == float64(int(val))
}
