package main

import "fmt"

func main() {

	fmt.Println("Hello world")

	var age int = 5

	var falso int = 10

	for i := 0; i < age; i++ {
		var v string
		if i == 0 {
			v = "ano!"
		} else {
			v = "anos!"
		}
		fmt.Println("Eu tenho", i+1, v)
	}

	var pt *int = &age

	println("age:", age, "pt:", *pt)
	age++
	println("age:", age, "pt:", *pt)
	*pt++
	println("age:", age, "pt:", *pt)

}
