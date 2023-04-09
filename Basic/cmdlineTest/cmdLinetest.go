package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	arg := os.Args[1:]
	var iArg = []int{}

	for _, i := range arg {
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArg = append(iArg, val)
	}
	max := 0
	for _, j := range iArg {
		if j > max {
			max = j
		}
	}
	println("max number is: ", max)
}
