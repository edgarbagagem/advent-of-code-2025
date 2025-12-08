package main

import (
	"edgar/aoc2025/utils"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <input_file>")
	}
	inputPath := os.Args[1]

	data, err := utils.InputStrings(inputPath)
	if err != nil {
		log.Fatalf("error parsing data from input: %s", err)
	}

	dial := 50
	zeroCount := 0
	for i, ins := range data {
		turn, err := strconv.Atoi(ins[1:])
		if err != nil {
			log.Fatalf("error parsing rotation instruction number on line %d: %q, error: %v", i+1, ins, err)
		}

		if ins[0] == 'L' {
			dial -= turn
		}

		if ins[0] == 'R' {
			dial += turn
		}

		if dial%100 == 0 {
			zeroCount++
		}
	}

	fmt.Println(zeroCount)
}
