/*
	Problem: https://adventofcode.com/2023/day/1
*/

package day1

import (
	"fmt"
	"log"
	"regexp"
	"unicode"

	"github.com/alpha2303/aoc2023/utils"
)

var (
	Numbers        = regexp.MustCompile(utils.Pattern)
	ReverseNumbers = regexp.MustCompile(utils.ReverseString(utils.Pattern))
)

func extractNumber(phrase string) int {
	var n int = len(phrase)
	var num int = 0
	var i int = 0

	for i < n {
		if unicode.IsDigit(rune(phrase[i])) {
			digit, err := utils.ToInt(string(phrase[i]), false)
			if err != nil {
				log.Fatal(err)
			}

			num += digit
			break
		}
		i++
	}

	var j int = n - 1
	for j >= i {
		if unicode.IsDigit(rune(phrase[j])) {
			digit, err := utils.ToInt(string(phrase[j]), false)
			if err != nil {
				log.Fatal(err)
			}

			num = (num * 10) + digit
			break
		}
		j--
	}

	return num
}

func extractNumber2(phrase string) int {
	firstDigit, err := utils.ToInt(Numbers.FindString(phrase), true)
	if err != nil {
		log.Fatal(err)
	}

	lastDigit, err := utils.ToInt(utils.ReverseString(ReverseNumbers.FindString(utils.ReverseString(phrase))), true)
	if err != nil {
		log.Fatal(err)
	}

	return (firstDigit * 10) + lastDigit
}

func Trebuchet(partNo int) {
	stringVector, err := utils.ReadInput("./inputs/day1_trebuchet.txt", 1000)
	if err != nil {
		log.Fatal(err)
	}

	var sum int = 0
	if partNo == 1 {
		for _, v := range stringVector {
			sum += extractNumber(v)
		}
	} else {
		for _, v := range stringVector {
			sum += extractNumber2(v)
		}
	}

	fmt.Printf("Result: %d\n", sum)
}
