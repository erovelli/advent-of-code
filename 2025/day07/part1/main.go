// Advent of Code 2025 Day 7 part 1
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	totalSplitCount := 0
	var prevMask []byte
	var currMask []byte
	n := 0

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

		if start := strings.IndexRune(line, 'S'); start != -1 {
			n = len(line)
			prevMask = make([]byte, n)
			prevMask[start] = 1
		}
		currMask = make([]byte, n)
		for i, c := range line {
			if c == '^' && prevMask[i] == 1 {
				// guards against splitting stream against edge
				if 0 < i {
					currMask[i-1] = 1
				}
				if i < len(line)-1 {
					currMask[i+1] = 1
				}
				currMask[i] = 0 // clear beam on split
				totalSplitCount++
			} else if prevMask[i] == 1 {
				currMask[i] = 1 // carry beam if no obstruction
			}
		}
		prevMask = currMask
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total amount of times the beam will be split ", totalSplitCount)
}
