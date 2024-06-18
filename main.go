package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/alpha2303/aoc2023/solutions/day1"
	"github.com/alpha2303/aoc2023/solutions/day2"
	"github.com/alpha2303/aoc2023/solutions/day3"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal(fmt.Errorf("arguments required: day_number, part_number"))
	}

	dayNo := os.Args[1]
	partNo, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	switch dayNo {
	case "1":
		day1.Trebuchet(partNo)
	case "2":
		day2.CubeConundrum(partNo)
	case "3":
		day3.GearRatios(partNo)
	default:
		log.Fatal(fmt.Errorf("invalid day number provided"))
	}
}
