package main

func Number(busStops [][2]int) int {
	total := 0
	for i := 0; i < len(busStops); i++ {
		total += busStops[i][0]
		total -= busStops[i][1]
	}
	return total
}
