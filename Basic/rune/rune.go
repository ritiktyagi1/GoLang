package main

import (
	"fmt"
	"unicode/utf8"
)

var p1 = fmt.Println

func main() {
	rStr := "abcdef"
	p1("Rune Count: ", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
}
