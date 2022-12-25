package main

func TowerBuilder(nFloors int) []string {
	tower := make([]string, nFloors)
	for i := 1; i <= nFloors; i++ {
		str := ""
		nStars := 2*i - 1
		for j := 0; j < nFloors-i; j++ {
			str += " "
		}
		for j := 0; j < nStars; j++ {
			str += "*"
		}
		for j := 0; j < nFloors-i; j++ {
			str += " "
		}
		tower = append(tower, str)
	}

	return tower
}
