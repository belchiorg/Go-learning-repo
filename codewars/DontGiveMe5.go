package main

func CheckContains5(n int) bool {

	for n != 0 {
		if n%10 == 5 || n%10 == -5 {
			return false
		}
		n /= 10
	}
	return true
}

func DontGiveMeFive(start int, end int) int {
	count := 0
	for i := start; i <= end; i++ {
		if CheckContains5(i) {
			count++
		}
	}
	return count
}
