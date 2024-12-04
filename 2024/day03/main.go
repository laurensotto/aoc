package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	answer1Chan := make(chan int)
	answer2Chan := make(chan int)
	time1Chan := make(chan int64)
	time2Chan := make(chan int64)

	go func() {
		start := time.Now()
		result := part1(input)
		duration := time.Since(start).Milliseconds()
		answer1Chan <- result
		time1Chan <- duration
	}()

	go func() {
		start := time.Now()
		result := part2(input)
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

func part1(input string) int {
	var calculationTotal = 0

	pattern := `mul\([0-9]{1,3},[0-9]{1,3}\)`

	re, _ := regexp.Compile(pattern)

	matches := re.FindAllString(input, -1)

	for _, match := range matches {
		pattern := `[0-9]{1,3}`

		re, _ := regexp.Compile(pattern)

		stringFormattedInts := re.FindAllString(match, -1)

		valueOne, _ := strconv.Atoi(stringFormattedInts[0])
		valueTwo, _ := strconv.Atoi(stringFormattedInts[1])

		calculationTotal += valueOne * valueTwo
	}

	return calculationTotal
}

func part2(input string) int {
	var calculationTotal = 0

	pattern := `mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`

	re, _ := regexp.Compile(pattern)

	matches := re.FindAllString(input, -1)

	do := true
	for _, match := range matches {
		if match == "do()" {
			do = true
			continue
		}

		if match == "don't()" {
			do = false
			continue
		}

		if !do {
			continue
		}

		pattern := `[0-9]{1,3}`

		re, _ := regexp.Compile(pattern)

		stringFormattedInts := re.FindAllString(match, -1)

		valueOne, _ := strconv.Atoi(stringFormattedInts[0])
		valueTwo, _ := strconv.Atoi(stringFormattedInts[1])

		calculationTotal += valueOne * valueTwo
	}

	return calculationTotal
}
