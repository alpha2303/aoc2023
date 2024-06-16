/*
	Problem: https://adventofcode.com/2023/day/2
*/

package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/alpha2303/aoc2023/utils"
)

var (
	ColorMap = map[string]int{
		"red":   0,
		"green": 1,
		"blue":  2,
	}
)

func getMaxColorCount(phrase string) ([3]int, error) {
	var maxGameColors [3]int = [3]int{0, 0, 0}

	gameInstances := strings.FieldsFunc(string(phrase[7:]), func(r rune) bool {
		return r == ',' || r == ';'
	})

	for i, instance := range gameInstances {
		countMap := strings.Split(instance, " ")
		// fmt.Println(countMap)
		// fmt.Println(countMap[1])
		count, err := strconv.Atoi(strings.Trim(countMap[1], " "))
		if err != nil {
			return maxGameColors, fmt.Errorf("integer conversion error at instance %d", (i + 1))
		}

		if maxGameColors[ColorMap[countMap[2]]] < count {
			maxGameColors[ColorMap[countMap[2]]] = count
		}
	}

	return maxGameColors, nil
}

func CubeConundrum(partNo int) {
	var colorLimit [3]int = [3]int{12, 13, 14}
	var sum int = 0
	inputVector, err := utils.ReadInput("./inputs/day2_cubeConundrum.txt", 100)
	if err != nil {
		log.Fatal(err)
	}

	for i, v := range inputVector {
		maxGameColors, err := getMaxColorCount(v)
		if err != nil {
			log.Fatal(err)
		}

		if partNo == 1 {
			var flag bool = true
			for j := 0; j < 3; j++ {
				if maxGameColors[j] > colorLimit[j] {
					flag = false
					break
				}
			}

			if flag {
				sum += (i + 1)
			}
		} else {
			sum += (maxGameColors[0] * maxGameColors[1] * maxGameColors[2])
		}

	}

	fmt.Printf("Day 2 Result: %d\n", sum)

}
