package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {

	for i := 0; i < 5; i++ {
		pl(i)
	}
	pl()
	x := 5
	for x > 0 {
		pl(x)
		x--
	}
	pl()
	aNums := []int{1, 4, 2, 6, 3}
	for _, num := range aNums {
		pl(num)
	}

}
