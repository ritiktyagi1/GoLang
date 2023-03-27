package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var p1 = fmt.Println

func main() {
	cV1 := 1.5
	cV2 := int(cV1)
	p1(cV2)
	// string to integer
	cV3 := "566059"
	cV4, err := strconv.Atoi(cV3)
	p1(cV4, err, reflect.TypeOf(cV4))
	// Convert string to float
	cV7 := "3.14"
	// Handling potential errors (Prints if err == nil)
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		p1(cV8)
	}
	// Use Sprintf to convert from float to string
	cV9 := fmt.Sprintf("%f", 3.14)
	p1(cV9)
}
