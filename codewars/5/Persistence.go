package main

func MultiplyAllNums(n int) int {
	total := 1
	for n > 0 {

		total *= n % 10
		n /= 10
	}
	return total
}

func Persistence(n int) int {
	count := 0
	for n >= 10 {
		n = MultiplyAllNums(n)
		count++
	}
	return count
}
