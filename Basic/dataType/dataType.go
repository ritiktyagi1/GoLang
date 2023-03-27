package main

import (
	"fmt"
	"reflect"
)

var p1 = fmt.Println

func main() {
	// int , float64 , bool, string, rune
	// Default type 0, 0.0, false, ""
	p1(reflect.TypeOf(35))
	p1(reflect.TypeOf(35.5))
	p1(reflect.TypeOf(true))
	p1(reflect.TypeOf("Heyy"))
	p1(reflect.TypeOf('🙂'))
}
