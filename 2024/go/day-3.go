package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Day3 struct{}

func (Day3) Part1() int {
	input := readTextFromInputFile(3)

	capturingPattern := `mul\((\d{1,3}),(\d{1,3})\)`
	capturingRe, err := regexp.Compile(capturingPattern)
	if err != nil {
		log.Fatalf("Invalid capturing regex pattern: %v", err)
	}

	submatches := capturingRe.FindAllStringSubmatch(input, -1)

	total := 0
	for _, submatch := range submatches {
		// submatch[0] is the entire match
		x, _ := strconv.Atoi(submatch[1])
		y, _ := strconv.Atoi(submatch[2])
		total += (x * y)
	}
	return total
}

func (Day3) Part2() int {
	input := readTextFromInputFile(3)

	// This flag will be toggled while sliding through the text.
	// don't() will turn it off.
	enabledFlag := true

	pattern := `(?:mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\))`
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatalf("Invalid capturing regex pattern: %v", err)
	}

	matches := re.FindAllString(input, -1)

	total := 0
	for _, match := range matches {
		switch match {
		case `do()`:
			enabledFlag = true
		case `don't()`:
			enabledFlag = false
		default:
			if enabledFlag {
				nums := strings.Split((match[4 : len(match)-1]), ",")
				x, _ := strconv.Atoi(nums[0])
				y, _ := strconv.Atoi(nums[1])
				total += (x * y)

			}
		}
		if match == `do()` {
			enabledFlag = true
		}
		if match == `don't()` {
			enabledFlag = false
		}

	}

	return total
}
