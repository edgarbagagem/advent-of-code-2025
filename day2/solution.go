package day2

import (
	"edgar/aoc2025/utils"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Solution(part int) error {
	switch part {
	case 1:
		return partOne()
	case 2:
	default:
		return fmt.Errorf("invalid part %d for day 2", part)
	}

	return nil
}

func partOne() error {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <input_file>")
	}
	inputPath := os.Args[1]

	data, err := utils.InputStrings(inputPath)
	if err != nil {
		log.Fatalf("error parsing data from input: %s", err)
	}

	if len(data) > 1 {
		log.Fatal("Error, expecting only a single line")
	}

	line := strings.Split(data[0], ",")

	// Example line -> 10-100,200-300
	sum := 0
	for _, valueRange := range line {
		pair := strings.Split(valueRange, "-")
		lower, err := strconv.Atoi(pair[0])
		if err != nil {
			return err
		}

		upper, err := strconv.Atoi(pair[1])
		if err != nil {
			return err
		}

		for i := lower; i <= upper; {
			digitsCount := len(strconv.Itoa(i))

			// isOdd
			if digitsCount&1 == 1 {
				nextEven := int(math.Pow10(digitsCount))
				if nextEven <= upper {
					i = nextEven
					continue
				}
				break
			}
			digitsCountHalf := digitsCount >> 1
			bigHalf := i / int(math.Pow10(digitsCountHalf))
			smallHalf := i % int(math.Pow10(digitsCountHalf))

			if bigHalf == smallHalf {
				sum += i
			}
			i++
		}
	}

	fmt.Println(sum)
	return nil
}
