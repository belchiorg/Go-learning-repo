package main

import "strings"

func toWeirdCase(str string) string {

	final := ""
	arr := strings.Split(str, " ")
	for j, s := range arr {
		for i, n := range s {
			if i%2 == 0 {
				//even case

				final += strings.ToUpper(string(n))
			} else {
				//odd case

				final += strings.ToLower(string(n))
			}
		}
		if j != len(arr)-1 {
			final += " "
		}
	}
	return final
}
