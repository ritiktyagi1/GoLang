package main

import (
	"fmt"
	"sync"
	"time"
)

var pl = fmt.Println

func main() {
	balance := 0
	var m sync.Mutex

	counter := func() {
		m.Lock()
		for i := 0; i < 10000000; i++ {
			balance = balance + 10
		}
		m.Unlock()
	}
	go counter()
	go counter()

	time.Sleep(time.Second)
	fmt.Printf("Balance : %v", balance)
}
