// Advent of Code 2025 Day 2 part 1
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	count := 0

	// fetch input file
	if len(os.Args) < 2 {
		log.Fatal("missing file input")
	}

	// open file
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//
	reader := bufio.NewReader(file)
	for {
		// delimit input sequence by comma
		slice, err := reader.ReadString(',')
		if err != nil {
			if err == io.EOF {
				break // exit loop at end of file
			}
			log.Fatal(err) // any other error is fatal
		}

		slice = strings.TrimRight(slice, ",") // remove trailing comma

		// process range
		parts := strings.Split(slice, "-")
		start, err1 := strconv.Atoi(parts[0][:])
		end, err2 := strconv.Atoi(parts[1][:])

		if err1 != nil || err2 != nil {
			panic("failed to parse numbers")
		}

		for i := start; i <= end; i++ {

			// check for duplicate sequence halves
			mid := len(strconv.Itoa(i)) / 2
			if strconv.Itoa(i)[:mid] == strconv.Itoa(i)[mid:] {
				count += i
			}
		}

	}
	fmt.Printf("Your puzzle answer is: %d\n", count)
}
