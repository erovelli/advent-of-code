// Advent of Code 2025 Day 3 part 2
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const JOLTAGE_WIDTH = 12

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
		maxLineJoltage, err := processLine(line)
		if err != nil {
			panic("failed to parse numbers")
		}
		maxJoltage += maxLineJoltage
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("your the total output joltage is: ", maxJoltage)
}

func processLine(line string) (int, error) {
	n := len(line)
	maxLineJoltage := ""

	t := 0

	// greedily pick the largest possible digit within the window of remaining positions.
	for i := range JOLTAGE_WIDTH {
		remainingWidth := n + i + 1 - JOLTAGE_WIDTH
		largest := "0"
		for j := t; j < remainingWidth; j++ {
			c := string(line[j])
			if c > largest {
				largest = c
				t = j + 1
			}
		}

		// last digit is the largest remaining but was skipped.
		if t == n {
			largest = string(line[n-1])
		}

		maxLineJoltage += largest
	}

	return strconv.Atoi(maxLineJoltage)
}
