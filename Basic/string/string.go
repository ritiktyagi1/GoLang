package main

import (
	"fmt"
	"strings"
)

var p1 = fmt.Println

func main() {
	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	p1(sV2)
	p1("Length: ", len(sV2))
	p1("Contains word: ", strings.Contains(sV2, "Word"))
	p1("o index: ", strings.Index(sV2, "o"))
	p1("Replace: ", strings.Replace(sV2, "o", "0", -1))

	sV3 := "\n Some sentneces \n"
	sV3 = strings.TrimSpace(sV3)
	p1("Split: ", strings.Split("a-b-c-d", "-"))
	p1("Lower: ", strings.ToLower(sV2))
	p1("Upper: ", strings.ToUpper(sV2))
	p1("Prefix: ", strings.HasPrefix("naman", "na"))
	p1("Suffix: ", strings.HasSuffix("naman", "an"))

}
