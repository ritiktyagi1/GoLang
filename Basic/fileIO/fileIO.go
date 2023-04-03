package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {
	//create a file
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// create list of primes

	iPrimeArr := []int{2, 3, 5, 7, 11}
	// create string array from int array

	var sPrrimeArr []string
	for _, i := range iPrimeArr {
		sPrrimeArr = append(sPrrimeArr, strconv.Itoa(i))
	}

	// Cycle through strings and write to file

	for _, num := range sPrrimeArr {
		_, err := f.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	// Open the created file
	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Read from file and print once per rline

	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		pl("Prime: ", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}

	// Check if file exists
	_, err = os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist) {
		pl("File Doesn't Exist")
	} else {
		f, err = os.OpenFile("data.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err := f.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}
	}
}
