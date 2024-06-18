/*
	Problem: https://adventofcode.com/2023/day/3
*/

package day3

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"unicode"

	"github.com/alpha2303/aoc2023/utils"
)

var notSymbols string = `[^.0-9]`
var notSymRegex *regexp.Regexp = regexp.MustCompile(notSymbols)
var offsets []int = []int{-1, 0, 1}

func idx2Str(row int, col int) string {
	return fmt.Sprintf("%d%d", row, col)
}

func isValidIdx(idx int, sliceSize int) bool {
	return idx >= 0 && idx < sliceSize
}

func getNumber(line string, row int, col int, visited *utils.Set[string], colCount int) int {
	var numString string = string(line[col])
	var offset int = 1

	for isValidIdx((col+offset), colCount) && unicode.IsDigit(rune(line[(col+offset)])) {
		numString = numString + string(line[(col+offset)])
		visited.Add(idx2Str(row, (col + offset)))
		offset++
	}

	offset = -1

	for isValidIdx((col+offset), colCount) && unicode.IsDigit(rune(line[(col+offset)])) {
		numString = string(line[(col+offset)]) + numString
		visited.Add(idx2Str(row, (col + offset)))
		offset--
	}

	// fmt.Printf("%s, %d, %d\n", line, row, col)

	number, err := strconv.Atoi(numString)
	if err != nil {
		log.Fatal(err)
	}

	return number

}

func checkAdjNums(inputVector []string, i int, j int, visited *utils.Set[string]) int {
	var subSum int = 0
	var rowCount int = len(inputVector)
	var colCount int = len(inputVector[i])
	for _, rowOffset := range offsets {
		row := i + rowOffset
		for _, colOffset := range offsets {
			col := j + colOffset

			if !isValidIdx(col, colCount) || !isValidIdx(row, rowCount) || visited.Has(idx2Str(row, col)) {
				continue
			}

			if unicode.IsDigit(rune(inputVector[i][j])) {
				fmt.Printf("%s, %d, %d\n", inputVector[row], row, col)
				subSum += getNumber(inputVector[row], row, col, visited, colCount)
			}
		}
	}

	return subSum
}

func GearRatios(partNo int) {
	inputVector, err := utils.ReadInput("./inputs/day3_gearRatios.txt", 140)
	if err != nil {
		log.Fatal(err)
	}

	var visited utils.Set[string]
	var sum int = 0
	for i, line := range inputVector {
		symIdxs := notSymRegex.FindAllStringIndex(line, len(line))
		if len(symIdxs) > 0 {
			for _, idx := range symIdxs {
				// fmt.Printf("%s, %d, %d\n", line, i, idx[0])
				sum += checkAdjNums(inputVector, i, idx[0], &visited)
			}
		}
	}

	fmt.Printf("Day 3 Result: %d\n", sum)
}
