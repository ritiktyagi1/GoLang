package main

import (
	"fmt"
)

var pl = fmt.Println

type rectangle struct {
	length, height float64
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}

// Go doesn't support inheritance, but it does
// support composition by embedding a struct
// in another

type contact struct {
	fName string
	lName string
	phone string
}
type business struct {
	name    string
	address string
	contact
}

func (b business) info() {
	fmt.Printf("Contact at %s is %s %s", b.name, b.contact.fName, b.contact.lName)
}
func main() {
	rect1 := rectangle{23.00, 45.00}
	pl("Area of rectangle: ", rect1.Area())
	con1 := contact{
		"James",
		"Wang",
		"555-1212",
	}

	bus1 := business{
		"ABC Plumbing",
		"234 North St",
		con1,
	}
	bus1.info()
}
