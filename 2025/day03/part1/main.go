// Advent of Code 2025 Day 3 part 1
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	maxJoltage := 0

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

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		lineLength := len(line)
		l, r := "0", "0"

		// construct largest two digit number from line
		for i, c := range line {
			c := string(c)
			if c > l && i < lineLength-1 {
				l = c
				r = string(line[i+1])
			} else if c > r {
				r = c
			}
		}
		lineMax, err := strconv.Atoi(l + r)
		if err != nil {
			panic("failed to parse numbers")
		}
		maxJoltage += lineMax

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("your the total output joltage is: ", maxJoltage)
}
