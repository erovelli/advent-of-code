// Advent of Code 2025 Day 6 part 1
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
	homeworkSum := 0
	var operands [][]string
	var operators []string

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
		fields := strings.Fields(scanner.Text())

		if fields[0] == "+" || fields[0] == "*" {
			operators = fields
		} else {
			operands = append(operands, fields)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(operands[0]) != len(operators) {
		log.Panic("operand list to operator mismatch")
	}

	for i, operator := range operators {
		result := 0
		if operator == "*" {
			result = 1
		}
		for operand := range operands {
			num, err := strconv.Atoi(operands[operand][i])
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
		homeworkSum += result

	}

	fmt.Println("Total output for the math homework  ", homeworkSum)
}
