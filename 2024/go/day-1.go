package main

import (
	"strconv"
	"strings"
)

type Day1 struct{}

func (Day1) Part1() int {
	output := 0
	list1, list2, inputLength := getSortedLists()

	for i := 0; i < inputLength; i++ {
		if list1[i] == list2[i] {
			continue
		} else if list1[i] > list2[i] {
			output += (list1[i] - list2[i])
			continue
		} else {
			output += (list2[i] - list1[i])
		}
	}

	return output
}

func (Day1) Part2() int {
	list1, list2, _ := getSortedLists()
	list2Map := make(map[int]int)
	output := 0

	for _, num := range list2 {
		list2Map[num]++
	}

	for _, num := range list1 {
		if _, ok := list2Map[num]; ok {
			output += (num * list2Map[num])
		}
	}

	return output
}

func getSortedLists() (left []int, right []int, length int) {
	lines := readLinesFromInputFile(1)
	list1, list2 := getListsFromLines(lines)

	quickSort(list1)
	quickSort(list2)

	return list1, list2, len(lines)
}

func getListsFromLines(lines []string) (left []int, right []int) {
	inputLength := len(lines)
	list1 := make([]int, inputLength)
	list2 := make([]int, inputLength)

	for i, line := range lines {
		if i >= 1000 {
			break
		}
		var num1, num2 int
		items := strings.Split(line, "   ")

		num1, err := strconv.Atoi(items[0])
		if err != nil {
			continue
		}
		list1[i] = num1

		num2, err = strconv.Atoi(items[1])
		if err != nil {
			continue
		}
		list2[i] = num2
	}

	return list1, list2
}

func quickSort(input []int) {
	if len(input) < 2 {
		return
	}

	pivotIndex := partition(input)

	quickSort(input[:pivotIndex])
	quickSort(input[pivotIndex+1:])
}

func partition(input []int) int {
	pivot := input[len(input)-1]
	i := -1

	for j := 0; j < len(input)-1; j++ {
		if input[j] < pivot {
			i++

			input[i], input[j] = input[j], input[i]
		}
	}

	input[i+1], input[len(input)-1] = input[len(input)-1], input[i+1]
	return i + 1
}
