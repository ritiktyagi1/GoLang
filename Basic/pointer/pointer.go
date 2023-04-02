package main

import (
	"fmt"
)

var pl = fmt.Println

func changeVal(x *int) {
	*x = 4
}

func changeArrVal(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] += 1
	}
}
func main() {
	a := 10
	var aPtr *int = &a
	pl("a address: ", aPtr)
	pl("a value: ", *aPtr)
	*aPtr = 5
	pl("a value before changing: ", a)
	changeVal(&a)
	pl("a value after change :", a)

	arr1 := [4]int{1, 2, 3, 4}
	changeArrVal(&arr1)
	pl(arr1)
}
