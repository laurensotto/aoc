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

	intLines := make([]int, len(lines))

	for i := 0; i < len(lines); i++ {
		intLines[i], _ = strconv.Atoi(lines[i])
	}

	answer1Chan := make(chan int)
	answer2Chan := make(chan int)
	time1Chan := make(chan int64)
	time2Chan := make(chan int64)

	go func() {
		start := time.Now()
		result := part1(intLines)
		duration := time.Since(start).Milliseconds()
		answer1Chan <- result
		time1Chan <- duration
	}()

	go func() {
		start := time.Now()
		result := part2(intLines)
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

func part1(list []int) int {
	totalIncreased := 0

	for i := 1; i < len(list); i++ {
		if list[i] > list[i-1] {
			totalIncreased++
		}
	}

	return totalIncreased
}

func part2(list []int) int {
	totalIncreased := 0

	prevSlider := list[0] + list[1] + list[2]
	for i := 1; i < len(list)-2; i++ {
		currentSlider := list[i] + list[i+1] + list[i+2]

		if currentSlider > prevSlider {
			totalIncreased++
		}

		prevSlider = currentSlider
	}

	return totalIncreased
}
