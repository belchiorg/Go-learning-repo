package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SimpleAssembler(program []string) map[string]int {
	regs := map[string]int{}

	for i := 0; i < len(program); i++ {
		instructions := strings.Split(program[i], " ")
		switch instructions[0] {
		case "mov":
			val, ok := regs[instructions[2]]
			if ok {
				regs[instructions[1]] = val
			} else {
				regs[instructions[1]], _ = strconv.Atoi(instructions[2])
			}
			break
		case "inc":
			regs[instructions[1]]++
			break
		case "dec":
			regs[instructions[1]]--
			break
		case "jnz":
			val, ok := regs[instructions[1]]
			if ok {
				if val != 0 {
					v, _ := strconv.Atoi(instructions[2])
					i--
					i += v
				}
			} else {
				v, _ := strconv.Atoi(instructions[1])
				if v != 0 {
					va, _ := strconv.Atoi(instructions[2])
					i--
					i += va
				}
			}
			break
		}
	}
	return regs
}

func main() {

	fmt.Println("Hello world")

	mapa := SimpleAssembler([]string{"mov d 100", "dec d", "mov b d", "jnz b -2", "inc d", "mov a d", "jnz 5 10", "mov c a"})
	for key, val := range mapa {
		println(key, val)
	}
}
