package main

import (
	"fmt"
	"log"
	"os"
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

	grid := make([][]string, len(lines))

	for i := range lines {
		gridRow := strings.Split(lines[i], "")

		grid[i] = gridRow
	}

	answer1Chan := make(chan int)
	answer2Chan := make(chan int)
	time1Chan := make(chan int64)
	time2Chan := make(chan int64)

	go func() {
		start := time.Now()
		result := part1(grid)
		duration := time.Since(start).Milliseconds()
		answer1Chan <- result
		time1Chan <- duration
	}()

	go func() {
		start := time.Now()
		result := part2(grid)
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

func part1(grid [][]string) int {
	currentX, currentY := findStart(grid)
	direction := "up"
	guardOnGrid := true

	totalVisitedFields := 0

	for guardOnGrid {
		if grid[currentY][currentX] != "X" {
			totalVisitedFields++
			grid[currentY][currentX] = "X"
		}

		currentX, currentY = getNextPosition(currentX, currentY, direction)

		if isOffGrid(currentX, currentY, grid) {
			guardOnGrid = false
			continue
		}

		if grid[currentY][currentX] == "#" {
			currentX, currentY = getPreviousPosition(currentX, currentY, direction)
			direction = rotate(direction)
		}
	}

	return totalVisitedFields
}

func part2(grid [][]string) int {
	return 0
}

func findStart(grid [][]string) (int, int) {
	for y, row := range grid {
		for x, value := range row {
			if value == "^" {
				return x, y
			}
		}
	}
	return 0, 0
}

func getNextPosition(currentX, currentY int, direction string) (int, int) {
	switch direction {
	case "up":
		return currentX, currentY - 1
	case "down":
		return currentX, currentY + 1
	case "left":
		return currentX - 1, currentY
	default:
		return currentX + 1, currentY
	}
}

func getPreviousPosition(currentX, currentY int, direction string) (int, int) {
	switch direction {
	case "up":
		return currentX, currentY + 1
	case "down":
		return currentX, currentY - 1
	case "left":
		return currentX + 1, currentY
	default:
		return currentX - 1, currentY
	}
}

func rotate(direction string) string {
	switch direction {
	case "up":
		return "right"
	case "down":
		return "left"
	case "left":
		return "up"
	default:
		return "down"
	}
}

func isOffGrid(currentX, currentY int, grid [][]string) bool {
	if currentX < 0 || currentY < 0 || currentX > len(grid[0])-1 || currentY > len(grid)-1 {
		return true
	}
	return false
}
