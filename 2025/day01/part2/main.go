// Advent of Code 2025 Day 1 part 2
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	dialRange := 100
	currentPosition := 50
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

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()

		direction := line[0]                    // R or L
		rotation, err := strconv.Atoi(line[1:]) // strip direction
		if err != nil {
			log.Fatal("rotation missing character", err)
		}

		// increment count on full rotations
		count += rotation / dialRange

		if direction == 'L' {
			rotation = -1 * rotation
		}
		// can ignore positive case

		// increment count when crossing 0, but not starting from 0
		if currentPosition != 0 && (currentPosition+rotation < 0 || currentPosition+rotation > dialRange) {
			count++
		}

		// dial rotation
		currentPosition = (currentPosition + rotation + dialRange) % dialRange

		// increment on 0 position
		if currentPosition == 0 {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total rotations to or past position 0:", count)
}
