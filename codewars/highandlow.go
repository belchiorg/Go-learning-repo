package main

import (
	"sort"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {

	arr := strings.Split(in, " ")

	var int_arr []int

	for _, n := range arr {
		nn, _ := strconv.Atoi(n)
		int_arr = append(int_arr, nn)
	}

	sort.Ints(int_arr)

	return strconv.Itoa(int_arr[len(int_arr)-1]) + " " + strconv.Itoa(int_arr[0])
}
