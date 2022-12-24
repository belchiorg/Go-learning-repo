package main

func FindMissingLetter(chars []rune) rune {
	for i := 1; i < len(chars); i++ {
		if int(chars[i]) == int(chars[i-1])+2 {
			return rune(int(chars[i]) - 1)
		}
	}
	return 'a'
}
