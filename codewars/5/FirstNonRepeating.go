package main

import "strings"

func FirstNonRepeating(str string) string {
	if str == "" {
		return ""
	}
	if len(str) == 1 {
		return string(str[0])
	}

	newString := ""
	firstChar := str[0]
	count := 0

	for _, c := range str {
		if strings.ToLower(string(firstChar)) == strings.ToLower(string(c)) {
			count++
		} else {
			newString += string(c)
		}
	}
	if count == 1 {
		return string(firstChar)
	}
	return FirstNonRepeating(newString)
}
