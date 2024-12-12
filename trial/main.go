package main

import "fmt"

func main() {

	allkeys := make(map[string]int)

	a, b := allkeys["naga"]
	fmt.Println(a, b)
	allkeys["naga"] = 123
	c, d := allkeys["naga"]

	fmt.Println(c, d)

}
