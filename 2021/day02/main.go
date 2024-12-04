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
	instructionLines := make([]string, len(lines))
	valueLines := make([]int, len(lines))

	for i, line := range lines {
		values := strings.Split(line, " ")

		instructionLines[i] = values[0]
		valueLines[i], _ = strconv.Atoi(values[1])
	}

	answer1Chan := make(chan int)
	answer2Chan := make(chan int)
	time1Chan := make(chan int64)
	time2Chan := make(chan int64)

	go func() {
		start := time.Now()
		result := part1(instructionLines, valueLines)
		duration := time.Since(start).Milliseconds()
		answer1Chan <- result
		time1Chan <- duration
	}()

	go func() {
		start := time.Now()
		result := part2(instructionLines, valueLines)
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

func part1(instructionLines []string, valueLines []int) int {
	horizontalPosition := 0
	verticalPosition := 0

	for i := 0; i < len(instructionLines); i++ {
		move := valueLines[i]

		switch instructionLines[i] {
		case "forward":
			horizontalPosition += move
		case "down":
			verticalPosition += move
		case "up":
			verticalPosition -= move
		}
	}

	return horizontalPosition * verticalPosition
}

func part2(instructionLines []string, valueLines []int) int {
	horizontalPosition := 0
	verticalPosition := 0
	aim := 0

	for i := 0; i < len(instructionLines); i++ {
		move := valueLines[i]

		switch instructionLines[i] {
		case "forward":
			horizontalPosition += move
			verticalPosition += aim * move
		case "down":
			aim += move
		case "up":
			aim -= move
		}
	}

	return horizontalPosition * verticalPosition
}
