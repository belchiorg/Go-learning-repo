package main

func MinMax(arr []int) [2]int {
	var min int = arr[0]
	var max int = arr[0]
	var leng = len(arr)

	for i := 1; i < leng; i++ {
		if min > arr[i] {
			min = arr[i]
		}
		if max < arr[i] {
			max = arr[i]
		}
	}
	return [2]int{min, max}
}
