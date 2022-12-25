package main

import (
	"sort"
	"strconv"
	"strings"
)

type Yau struct {
	weight string
	val    int
}

func OrderWeight(strng string) string {
	arr := strings.Split(strng, " ")
	weightvals := []Yau{}
	for _, n := range arr {
		sum := 0
		nums := strings.Split(n, "")
		for _, num := range nums {
			m, _ := strconv.Atoi(num)
			sum += m
		}
		weightvals = append(weightvals, Yau{weight: n, val: sum})
	}

	sort.Slice(weightvals, func(i, j int) bool {
		if weightvals[i].val < weightvals[j].val {
			return true
		} else if weightvals[i].val == weightvals[j].val {
			return strings.Compare(weightvals[i].weight, weightvals[j].weight) < 0
		}
		return false
	})

	final := ""

	for _, n := range weightvals {
		final += n.weight + " "
	}

	return strings.TrimRight(final, " ")
}
