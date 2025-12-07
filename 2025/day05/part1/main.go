// Advent of Code 2025 Day 5 part 1
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var freshRange [][]int
	freshIngredients := 0

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
		fmt.Println(line)

		// store range
		if strings.Contains(line, "-") {

			// process range
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				panic("range must only have a start and end <val1>-<val2>")
			}

			start, err1 := strconv.Atoi(parts[0])
			end, err2 := strconv.Atoi(parts[1])

			if err1 != nil || err2 != nil {
				panic("failed to parse numbers")
			}

			freshRange = append(freshRange, []int{start, end})

		} else if len(line) > 0 {
			// check if ingredient id is in fresh range
			for _, v := range freshRange {
				ingredient, err := strconv.Atoi(line)
				if err != nil {
					panic("ingredient string to int conversion failure")
				}
				if v[0] <= ingredient && ingredient <= v[1] {
					freshIngredients++
					break
				}
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total number of fresh ingredient:  ", freshIngredients)
}
