package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	sl1 := make([]string, 4)

	sl1[0] = "Hajimemashite"
	sl1[1] = "Watashi"
	sl1[2] = "Ritik"
	sl1[3] = "desu"

	pl("Size of slize: ", len(sl1))

	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}

	for _, x := range sl1 {
		pl(x)
	}

	sl2 := []int{21, 34, 11}
	pl(sl2)

	sArr := [5]int{1, 2, 3, 4, 5}

	sl3 := sArr[:3]
	pl(sl3)

	pl("1st 3  :", sArr[:3])
	pl("Last 3 :", sArr[2:])

	sArr[1] = 30
	pl("sl3: ", sl3)

	sl3[0] = 1
	pl("sArr: ", sArr)
	// Append a value to a slice (Also overwrites array)
	sl3 = append(sl3, 12)
	pl("sl3 :", sl3)
	pl("sArr :", sArr)

	// Printing empty slices will return nils which show
	// as empty slices
	sl4 := make([]string, 6)
	pl("sl4 :", sl4)
	pl("sl4[0] :", sl4[0])
}
