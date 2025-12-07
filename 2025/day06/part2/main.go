// Advent of Code 2025 Day 6 part 2
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// test input:
// 123 328  51 64
//  45 64  387 23
//   6 98  215 314
// *   +   *   +
/*
The rightmost problem is 4 + 431 + 623 = 1058
The second problem from the right is 175 * 581 * 32 = 3253600
The third problem from the right is 8 + 248 + 369 = 625
Finally, the leftmost problem is 356 * 24 * 1 = 8544
Now, the grand total is 1058 + 3253600 + 625 + 8544 = 3263827.

*/
func main() {
	homeworkSum := 0
	var operators []string
	var grid []string
	var fields [][]string

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
		lineFields := strings.Fields(line)

		// problem operators
		if lineFields[0] == "+" || lineFields[0] == "*" {
			operators = lineFields
		} else {
			// problem operands
			fields = append(fields, lineFields) // whitespace removed
			grid = append(grid, line)           // raw problem grid
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	problemWidths := maxProblemWidths(fields)

	problems := spaceAwareProblemSet(grid, fields, problemWidths)

	if len(problems) != len(operators) {
		log.Panic("operand list to operator mismatch")
	}

	// solve each problem and return results sum
	for i, operator := range operators {
		homeworkSum += cephalopodMathSolver(problems[i], problemWidths[i], operator)
	}

	fmt.Println("Total output for the math homework  ", homeworkSum)
}

func cephalopodMathSolver(problem []string, width int, operator string) int {
	result := 0
	if operator == "*" {
		result = 1
	}
	for i := range width {
		operand := ""
		// operands are constructed by concatenating like param indexes
		for _, param := range problem {
			// skip whitespace
			if param[i] != ' ' {
				operand += string(param[i])
			}
		}
		num, err := strconv.Atoi(operand)
		if err != nil {
			log.Panic("could not convert string to int")
		}
		switch operator {
		case "+":
			result += num
		case "*":
			result *= num
		}
	}
	return result
}

func maxProblemWidths(fields [][]string) []int {
	m := len(fields)
	n := len(fields[0])
	problemWidths := make([]int, n)
	for col := range n {
		maxWidth := 0
		for row := range m {
			maxWidth = max(maxWidth, len(fields[row][col]))
		}
		problemWidths[col] = maxWidth

	}
	return problemWidths

}

func spaceAwareProblemSet(grid []string, fields [][]string, problemWidths []int) [][]string {
	m := len(fields)
	n := len(fields[0])
	problems := make([][]string, n)
	start := 0

	// problem
	for col := range n {
		problems[col] = make([]string, m)
		end := start + problemWidths[col]

		// opperand per problem
		for row := range m {
			problems[col][row] = grid[row][start:end]
		}
		// begin next column after gutter whitespace
		start += problemWidths[col] + 1
	}

	return problems
}
