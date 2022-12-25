package main

func DigitalRoot(n int) int {
	sum := 0
	for n >= 10 {
		sum += n % 10
		n /= 10
	}
	sum += n
	if sum >= 10 {
		return DigitalRoot(sum)
	}
	return sum
}
