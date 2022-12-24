package main

func FindOdd(seq []int) int {
	var seen []int
	for i := 0; i < len(seq); i++ {
		for _, r := range seen {
			if seq[i] == r {
				continue
			}
		}

		seen = append(seen, seq[i])

		reps := 1

		j := i + 1
		for seq[j] == seq[i] {
			reps++
			j++
		}
		if reps%2 == 1 {
			return seq[i]
		}

		i = j - 1
	}

	return 0
}

func FindOddXOR(seq []int) int {
	res := 0

	for _, n := range seq {
		res ^= n //Usa esta merda de XOR operator, basicamente vai ver os numeros em binario, cancelando aqueles que são iguais :D
	}

	return res
}
