package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	filename := "challenge.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	answer1, time1, answer2, time2 := solve(string(data))
	fmt.Printf("Part 1: %d (Time: %d ms)\n", answer1, time1)
	fmt.Printf("Part 2: %d (Time: %d ms)\n", answer2, time2)
}

func solve(input string) (int, int64, int, int64) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var manualPages [][]int
	manualRules := make(map[int][]int)

	foundBreakPoint := false
	for i := range lines {
		if lines[i] == "" {
			foundBreakPoint = true
			continue
		}

		if !foundBreakPoint {
			values := strings.Split(lines[i], "|")

			beforeValue, _ := strconv.Atoi(values[0])
			afterValue, _ := strconv.Atoi(values[1])

			if _, ok := manualRules[beforeValue]; !ok {
				manualRules[beforeValue] = []int{}
			}

			manualRules[beforeValue] = append(manualRules[beforeValue], afterValue)

			continue
		}

		values := strings.Split(lines[i], ",")
		intValues := make([]int, len(values))
		for i := range values {
			intValue, _ := strconv.Atoi(values[i])

			intValues[i] = intValue
		}

		manualPages = append(manualPages, intValues)
	}

	answer1Chan := make(chan int)
	answer2Chan := make(chan int)
	time1Chan := make(chan int64)
	time2Chan := make(chan int64)

	go func() {
		start := time.Now()
		result := part1(manualRules, manualPages)
		duration := time.Since(start).Milliseconds()
		answer1Chan <- result
		time1Chan <- duration
	}()

	go func() {
		start := time.Now()
		result := part2(manualRules, manualPages)
		duration := time.Since(start).Milliseconds()
		answer2Chan <- result
		time2Chan <- duration
	}()

	part1Result := <-answer1Chan
	time1Result := <-time1Chan
	part2Result := <-answer2Chan
	time2Result := <-time2Chan

	return part1Result, time1Result, part2Result, time2Result
}

func part1(manualRules map[int][]int, manualPages [][]int) int {
	totalMiddlePages := 0

	for i := range manualPages {
		if isPageValid(manualRules, manualPages[i]) {
			totalMiddlePages += manualPages[i][len(manualPages[i])/2]
		}
	}

	return totalMiddlePages
}

func part2(manualRules map[int][]int, manualPages [][]int) int {
	totalMiddlePages := 0

	for i := range manualPages {
		if !isPageValid(manualRules, manualPages[i]) {
			sortedManualPage := sortManualPage(manualRules, manualPages[i])
			totalMiddlePages += sortedManualPage[len(sortedManualPage)/2]
		}
	}

	return totalMiddlePages
}

func isPageValid(manualRules map[int][]int, manualPage []int) bool {
	for i := 1; i < len(manualPage); i++ {
		if rule, ok := manualRules[manualPage[i]]; ok {

			for j, value := range manualPage {
				if j < i && contains(rule, value) {

					return false
				}

				if j > i && !contains(rule, value) {
					return false
				}
			}
		}
	}

	return true
}

func contains(slice []int, int int) bool {
	for _, value := range slice {
		if value == int {
			return true
		}
	}

	return false
}

func sortManualPage(manualRules map[int][]int, manualPage []int) []int {
	sortedManualPage := []int{manualPage[0]}

	for i := 1; i < len(manualPage); i++ {
		valueToPlace := manualPage[i]

		if rule, ok := manualRules[valueToPlace]; ok {
			var err error

			sortedManualPage, err = tryInsertForward(rule, sortedManualPage, valueToPlace)

			if err == nil {
				continue
			}
		}

		for j := len(sortedManualPage) - 1; j >= 0; j-- {
			if rule, ok := manualRules[manualPage[j]]; ok {
				if contains(rule, valueToPlace) {
					if j == 0 {
						sortedManualPage = insertBetween(sortedManualPage, 0, 1, valueToPlace)
						break
					}

					if j == len(sortedManualPage)-1 {
						sortedManualPage = append(sortedManualPage, valueToPlace)
						break
					}

					sortedManualPage = insertBetween(sortedManualPage, i-1, i, valueToPlace)
					break
				}
			}
		}

		if !contains(sortedManualPage, valueToPlace) {
			sortedManualPage = append(sortedManualPage, valueToPlace)
		}
	}
	return sortedManualPage
}

func tryInsertForward(rule []int, sortedManualPage []int, valueToPlace int) ([]int, error) {
	for j, value := range sortedManualPage {
		if contains(rule, value) {
			if j == 0 {
				sortedManualPage = append([]int{valueToPlace}, sortedManualPage...)
				return sortedManualPage, nil
			}

			if j == len(sortedManualPage)-1 {
				sortedManualPage = insertBetween(sortedManualPage, len(sortedManualPage)-2, len(sortedManualPage)-1, valueToPlace)
				return sortedManualPage, nil
			}

			sortedManualPage = insertBetween(sortedManualPage, j-1, j, valueToPlace)
			return sortedManualPage, nil
		}
	}

	return sortedManualPage, errors.New("not able to place value")
}

func insertBetween(array []int, index1 int, index2 int, value int) []int {
	if index2 != index1+1 {
		panic("Indexes are not neighbors")
	}

	newArray := make([]int, 0, len(array)+1)
	newArray = append(newArray, array[:index2]...)
	newArray = append(newArray, value)
	newArray = append(newArray, array[index2:]...)

	return newArray
}
