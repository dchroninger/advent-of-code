package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	Day1{}.Part1()
	Day1{}.Part2()
	Day2{}.Part1()
	Day2{}.Part2()
	Day3{}.Part1()
	fmt.Printf("Day 3 - Part 2 : %v\n", Day3{}.Part2())
}

func readLinesFromInputFile(day int) []string {
	filePath := fmt.Sprintf("../input/day-%d.txt", day)

	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read file %s: %v", filePath, err))
	}

	lines := strings.Split(string(data), "\n")

	return lines
}

func readTextFromInputFile(day int) string {
	filePath := fmt.Sprintf("../input/day-%d.txt", day)

	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read file %s: %v", filePath, err))
	}

	return string(data)
}
