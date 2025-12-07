// Advent of Code 2025 Day 5 part 2
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// custom type for 2d range sort
type ByFirst [][]int

// Implement sort.Interface
func (a ByFirst) Len() int           { return len(a) }
func (a ByFirst) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFirst) Less(i, j int) bool { return a[i][0] < a[j][0] }

func main() {
	var freshRange [][]int
	totalFreshIngredientIds := 0

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

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sort.Sort(ByFirst(freshRange))
	start, end := 0, 0
	for _, v := range freshRange {

		// new range is a subset of existing range
		if start <= v[0] && v[1] <= end {
			continue
		}

		// partial range overlap - add the difference to total
		if v[0] <= end && end < v[1] {
			totalFreshIngredientIds += v[1] - end
			end = v[1]
		} else {
			// overwrite current range as there is no overlap
			start, end = v[0], v[1]
			totalFreshIngredientIds += end - start + 1 // add one as range is inclusive
		}
	}

	fmt.Println("Total number of fresh ingredient IDs:  ", totalFreshIngredientIds)
}
