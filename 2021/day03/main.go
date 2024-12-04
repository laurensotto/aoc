package main

import (
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

	splitLines := make([][]string, len(lines))

	for i, line := range lines {
		splitLines[i] = strings.Split(line, "")
	}

	answer1Chan := make(chan int)
	answer2Chan := make(chan int)
	time1Chan := make(chan int64)
	time2Chan := make(chan int64)

	go func() {
		start := time.Now()
		result := part1(splitLines)
		duration := time.Since(start).Milliseconds()
		answer1Chan <- result
		time1Chan <- duration
	}()

	go func() {
		start := time.Now()
		result := part2(splitLines)
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

func part1(lines [][]string) int {
	binaryTrueCount := make([]int, len(lines[0]))

	for _, line := range lines {
		for i, value := range line {
			if value == "1" {
				binaryTrueCount[i]++
			}
		}
	}

	binaryGamma := ""
	binaryEpsilon := ""

	for _, value := range binaryTrueCount {
		if value > len(lines)/2 {
			binaryGamma += "1"
			binaryEpsilon += "0"
		} else {
			binaryGamma += "0"
			binaryEpsilon += "1"
		}
	}

	decimalGamma, _ := strconv.ParseInt(binaryGamma, 2, 64)
	decimalEpsilon, _ := strconv.ParseInt(binaryEpsilon, 2, 64)

	return int(decimalGamma * decimalEpsilon)
}

func part2(lines [][]string) int {
	oxygenValue, _ := getValue(lines, 0, true)
	co2Value, _ := getValue(lines, 0, false)

	return int(oxygenValue * co2Value)
}

func getValue(lines [][]string, index int, pickLongerArray bool) (int64, error) {
	var binaryTrueArray [][]string
	var binaryFalseArray [][]string

	if len(lines) == 1 {
		oxygenBinary := ""

		for _, value := range lines[0] {
			oxygenBinary += value
		}

		return strconv.ParseInt(oxygenBinary, 2, 64)
	}

	for _, line := range lines {
		if line[index] == "1" {
			binaryTrueArray = append(binaryTrueArray, line)
		} else {
			binaryFalseArray = append(binaryFalseArray, line)
		}
	}

	if len(binaryFalseArray) > len(binaryTrueArray) {
		if pickLongerArray {
			return getValue(binaryFalseArray, index+1, true)
		} else {
			return getValue(binaryTrueArray, index+1, false)
		}
	}

	if pickLongerArray {
		return getValue(binaryTrueArray, index+1, true)
	} else {
		return getValue(binaryFalseArray, index+1, false)
	}
}
