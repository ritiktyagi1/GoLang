package main

import (
	"fmt"
	"math/rand"
	"time"
)

var p1 = fmt.Println

func main() {
	now := time.Now()
	p1(now.Year(), now.Month(), now.Day())
	p1(now.Hour(), now.Minute(), now.Second())
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	p1("Random :", randNum)
}
