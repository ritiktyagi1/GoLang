package main

import (
	"fmt"
)

var pl = fmt.Println

type customer struct {
	name    string
	address string
	bal     float64
}

func getCustInfo(c customer) {
	fmt.Printf("%s ows us %.2f\n", c.name, c.bal)
}

func newCustAdd(c *customer, address string) {
	c.address = address
}

// func main() {
// 	var c1 customer
// 	c1.name = "Ritik"
// 	c1.address = "Delhi"
// 	c1.bal = 455.56
// 	getCustInfo(c1)
// 	newCustAdd(&c1, "Tokyo")
// 	pl("Address :", c1.address)
// 	c2 := customer{"naman", "Delhi", 560.45}
// 	pl("name: ", c2.name)
// }
