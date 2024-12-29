package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Day2 struct{}

func (Day2) Part1() int {
	lines := readLinesFromInputFile(2)
	var safeCount, unsafeCount int

	for _, line := range lines {
		levels := strings.Split(line, " ")

		// the data parsers giving back a trailing row.
		if len(levels) == 1 {
			continue
		}

		var isIncrementing bool
		isSafe := true

		for i := 0; i < len(levels)-1; i++ {
			curr, err := strconv.Atoi(levels[i])
			if err != nil {
				panic(err.Error())
			}

			next, err := strconv.Atoi(levels[i+1])
			if err != nil {
				panic(err.Error())
			}

			// Set baseline on first iteration
			if i == 0 {
				isIncrementing = curr < next
			} else {
				// if the state of isIncrementing has changed, make it unsafe
				if isIncrementing != (curr < next) {
					isSafe = false
				}
			}

			diff := curr - next
			if diff < 0 {
				diff = -diff
			}

			if diff > 3 || diff == 0 {
				isSafe = false
			}
		}

		if isSafe {
			safeCount++
		} else {
			unsafeCount++
		}
	}

	return safeCount
}

func (Day2) Part2() int {
	lines := readLinesFromInputFile(2)
	safeCount := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			// skip any empty or trailing lines
			continue
		}

		// Parse line into integers
		strLevels := strings.Fields(line)
		levels := make([]int, len(strLevels))
		for i, s := range strLevels {
			num, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("Error converting to int:", err)
				continue
			}
			levels[i] = num
		}

		// Check safety under Part 2 rules
		if isPart2Safe(levels) {
			safeCount++
		}
	}

	return safeCount
}

// checkPart1Safe returns true if `levels` is strictly increasing
// *or* strictly decreasing by steps of 1..3 (no flats, no bigger jumps).
func checkPart1Safe(levels []int) bool {
	n := len(levels)
	if n < 2 {
		// 0 or 1 item is always trivially safe
		return true
	}

	// Check if strictly increasing
	increasing := true
	for i := 0; i < n-1; i++ {
		diff := levels[i+1] - levels[i]
		if diff <= 0 || diff > 3 {
			increasing = false
			break
		}
	}
	if increasing {
		return true
	}

	// Check if strictly decreasing
	decreasing := true
	for i := 0; i < n-1; i++ {
		diff := levels[i] - levels[i+1]
		if diff <= 0 || diff > 3 {
			decreasing = false
			break
		}
	}
	return decreasing
}

// isPart2Safe returns true if `levels` is already Part-1-safe
// OR if removing exactly one element makes it Part-1-safe.
func isPart2Safe(levels []int) bool {
	// If already safe by Part 1 rules:
	if checkPart1Safe(levels) {
		return true
	}

	// Otherwise, try removing each index in turn
	for i := 0; i < len(levels); i++ {
		// Construct a slice skipping index i
		newLevels := make([]int, 0, len(levels)-1)
		newLevels = append(newLevels, levels[:i]...)
		newLevels = append(newLevels, levels[i+1:]...)

		if checkPart1Safe(newLevels) {
			return true
		}
	}

	return false
}
