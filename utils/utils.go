package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var (
	Pattern = "one|two|three|four|five|six|seven|eight|nine|1|2|3|4|5|6|7|8|9"
	Digits  = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
)

func ReadInput(path string, inputSize int) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	stringVector := make([]string, inputSize)
	i := 0
	for scanner.Scan() {
		stringVector[i] = scanner.Text()
		i++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stringVector, nil
}

func ReverseString(s string) string {
	var result string
	for _, c := range s {
		result = string(c) + result
	}
	return result
}

func ToInt(phrase string, includeString bool) (int, error) {
	if len(phrase) == 1 && unicode.IsDigit(rune(phrase[0])) {
		return strconv.Atoi(phrase)
	}

	if includeString {
		digit, ok := Digits[phrase]
		if !ok {
			return -1, fmt.Errorf("unable to find digit for phrase: %s", phrase)
		}

		return digit, nil
	}

	return -1, fmt.Errorf("input string is neither a digit nor the name of a digit: %s", phrase)
}
