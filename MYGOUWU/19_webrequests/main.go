package main

// https://pkg.go.dev/net/http
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://example.com"

func main() {
	fmt.Println("Web Request")

	resp, err := http.Get(url)

	CheckNilErr(err)

	fmt.Printf("Response is of type: %T\n", resp)

	defer resp.Body.Close()

	databitos, err := ioutil.ReadAll(resp.Body)

	CheckNilErr(err)

	fmt.Println(string(databitos))
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
