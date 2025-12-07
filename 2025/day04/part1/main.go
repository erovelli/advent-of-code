// Advent of Code 2025 Day 4 part 1
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var grid []string
var m int
var n int

const roll = '@'
const maxRolls = 4
const depth = 1
const perimeterWidth = (2 * depth) + 1

func main() {
	removableRolls := 0

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
		grid = append(grid, line)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	m, n = len(grid), len(grid[0])

	// Determine # of removable rolls
	for row := range m {
		for col := range n {
			if grid[row][col] == roll && isRollRemovable(row, col) {
				removableRolls++
			}
		}
	}
	fmt.Println("Number of removable rolls:  ", removableRolls)
}

func isRollRemovable(row int, col int) bool {
	rolls := 0
	for i := range perimeterWidth {
		x := row - 1 + i
		// guard if row is out of bounds
		if !(0 <= x && x < m) {
			continue
		}
		for j := range perimeterWidth {
			y := col - 1 + j
			// guard if column is out of bounds
			if !(0 <= y && y < n) {
				continue
			}

			// skip the original roll
			if x == row && y == col {
				continue
			}
			if grid[x][y] == roll {
				rolls++
				if rolls == maxRolls {
					return false
				}

			}
		}
	}
	return true
}
