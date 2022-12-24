package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 1}
	tes := 0
	tes ^= arr[5]
	tes ^= arr[0]
	fmt.Println("valor de tes^=arr", tes)
}
