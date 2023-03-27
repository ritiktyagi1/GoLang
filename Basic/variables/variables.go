package main

import (
	"fmt"
)

var p1 = fmt.Println

func main() {
	// var name type
	var vNmae string = "Ritik"
	var v1, v2 = 2.3, 4
	var v3 = "Hello"
	v4 := 2.4
	v4 = 5.6
	p1(vNmae, v1, v2, v3, v4)
}
