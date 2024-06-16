package utils

import (
	"bufio"
	"os"
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
