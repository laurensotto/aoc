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

	var reports [][]int

	for _, line := range lines {
		values := strings.Split(line, " ")
		report := make([]int, len(values))
		for i, value := range values {
			intValue, _ := strconv.Atoi(value)
			report[i] = intValue
		}
		reports = append(reports, report)
	}

	answer1Chan := make(chan int)
	answer2Chan := make(chan int)
	time1Chan := make(chan int64)
	time2Chan := make(chan int64)

	go func() {
		start := time.Now()
		result := part1(reports)
		duration := time.Since(start).Milliseconds()
		answer1Chan <- result
		time1Chan <- duration
	}()

	go func() {
		start := time.Now()
		result := part2(reports)
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

func part1(reports [][]int) int {
	var validReports = 0
	for _, report := range reports {

		if checkReportValidity(report) {
			validReports++
		}
	}

	return validReports
}

func part2(reports [][]int) int {
	var validReports = 0
	for _, report := range reports {
		if checkReportValidity(report) {
			validReports++
			continue
		}

		for i := 0; i < len(report); i++ {
			var newReport []int

			for j := 0; j < len(report); j++ {
				if j != i {
					newReport = append(newReport, report[j])
				}
			}

			if checkReportValidity(newReport) {
				validReports++
				break
			}
		}
	}

	return validReports
}

func isNextIntGraduallyHigher(int1 int, int2 int) bool {
	if int1 <= int2 {
		return false
	}

	if int1-int2 > 3 {
		return false
	}

	return true
}

func checkReportValidity(report []int) bool {
	var isReportIncreasing = report[0] < report[1]

	for i, level := range report {
		if i == len(report)-1 {
			break
		}

		if isReportIncreasing && !isNextIntGraduallyHigher(report[i+1], level) {
			return false
		}

		if !isReportIncreasing && !isNextIntGraduallyHigher(level, report[i+1]) {
			return false
		}
	}

	return true
}
