package main

import (
	"fmt"
)

var p1 = fmt.Println

func main() {
	iAge := 8

	if (iAge >= 1) && (iAge <= 18) {
		p1("Important birthday")
	} else if (iAge == 21) || (iAge == 50) {
		p1("Important birthday")
	} else {
		p1("Not important")
	}
	p1("!True", !true)
}
