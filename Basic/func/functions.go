package main

import (
	"fmt"
)

var pl = fmt.Println

func hello() {
	pl("Hello Good morning")
}

func getOne(x int) int {
	return x + 1
}

func getTwo(x int) (int, int) {
	return x + 1, x + 2
}

func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You can't divide by zero")
	} else {
		return x / y, nil
	}
}

// Varadic Functions
func getSum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// passing array as a parameter
func getArrSum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}
func main() {
	hello()
	pl(getOne(2))
	pl(getTwo(3))
	pl(getQuotient(8, 0))
	pl(getQuotient(8, 2))
	pl(getSum(1, 2, 3, 4, 5))
	vArr := []int{23, 34, 50}
	pl(getArrSum(vArr))
}
