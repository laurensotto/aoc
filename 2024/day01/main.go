package main

import (
	"fmt"
	"log"
	"os"
	"slices"
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

	listOne := make([]int, len(lines))
	listTwo := make([]int, len(lines))

	for i := 0; i < len(lines); i++ {
		values := strings.Split(lines[i], "   ")

		listOne[i], _ = strconv.Atoi(values[0])
		listTwo[i], _ = strconv.Atoi(values[1])
	}

	answer1Chan := make(chan int)
	answer2Chan := make(chan int)
	time1Chan := make(chan int64)
	time2Chan := make(chan int64)

	go func() {
		start := time.Now()
		result := part1(listOne, listTwo)
		duration := time.Since(start).Milliseconds()
		answer1Chan <- result
		time1Chan <- duration
	}()

	go func() {
		start := time.Now()
		result := part2(listOne, listTwo)
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

func part1(listOne []int, listTwo []int) int {
	slices.Sort(listOne)
	slices.Sort(listTwo)

	totalDifference := 0
	for i := 0; i < len(listOne); i++ {
		totalDifference += getDifference(listOne[i], listTwo[i])
	}

	return totalDifference
}

func part2(listOne []int, listTwo []int) int {
	totalSimilarity := 0

	for _, value := range listOne {
		totalSimilarity += getSimilarityScore(value, listTwo)
	}

	return totalSimilarity
}

func getDifference(int1 int, int2 int) int {
	if int1 > int2 {
		return int1 - int2
	}

	return int2 - int1
}

func getSimilarityScore(int int, values []int) int {
	occurrence := 0

	for _, value := range values {
		if int == value {
			occurrence++
		}
	}

	return int * occurrence
}
