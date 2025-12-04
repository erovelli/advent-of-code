// Advent of Code 2025 Day 2 part 2
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
	invalidIdSum := 0
	slice := ""

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

	reader := bufio.NewReader(file)

	for {
		// delimit input sequence by comma
		slice, err = reader.ReadString(',')
		if err != nil {
			if err == io.EOF {
				break // exit loop at end of file
			}
			log.Fatal(err) // any other error is fatal
		}

		slice = strings.TrimRight(slice, ",") // remove trailing comma

		invalidIdSum += processIdRange(slice)

	}
	invalidIdSum += processIdRange(slice) // process last range after EOF
	fmt.Printf("Your puzzle answer is: %d\n", invalidIdSum)
}

func processIdRange(slice string) int {
	invalidIdSum := 0

	// process range
	parts := strings.Split(slice, "-")
	start, err1 := strconv.Atoi(parts[0][:])
	end, err2 := strconv.Atoi(parts[1][:])

	if err1 != nil || err2 != nil {
		panic("failed to parse numbers")
	}

	for id := start; id <= end; id++ {
		idStr := strconv.Itoa(id)
		for i := 0; i < len(idStr); i++ {
			// check for repeated substring pattern
			subStrOccurrence := strings.Count(idStr, idStr[:i])
			if subStrOccurrence > 1 && idStr == strings.Repeat(idStr[:i], subStrOccurrence) {
				invalidIdSum += id
				break
			}
		}
	}
	return invalidIdSum
}
