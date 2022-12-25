package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in go")
	content := "This needs to go in a file - Y@u"

	file, err := os.Create("./teste1.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println(length)
	defer file.Close()
	readFile("./teste1.txt")
}

func readFile(filename string) {
	databitos, err := ioutil.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databitos))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
