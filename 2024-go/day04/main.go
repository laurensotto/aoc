package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
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

	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := make([][]string, len(lines))

	for i := range lines {
		gridRow := strings.Split(lines[i], "")

		grid[i] = gridRow
	}

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
	var wg sync.WaitGroup
	resultChan := make(chan int, len(grid)*len(grid[0]))

	for y, row := range grid {
		for x, cell := range row {
			if cell == "X" {
				wg.Add(1)
				findXmas(x, y, grid, &wg, resultChan)
			}
		}
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	totalXmas := 0
	for result := range resultChan {
		totalXmas += result
	}

	return totalXmas
}

func part2(grid [][]string) int {
	var wg sync.WaitGroup
	resultChan := make(chan int, len(grid)*len(grid[0]))

	for y, row := range grid {
		for x, cell := range row {
			if cell == "A" {
				wg.Add(1)
				findCrossmas(x, y, grid, &wg, resultChan)
			}
		}
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	totalXmas := 0
	for result := range resultChan {
		totalXmas += result
	}

	return totalXmas
}

func findXmas(startX int, startY int, grid [][]string, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()

	needToCheckLeft := startX > 2
	needToCheckRight := startX < len(grid[0])-3
	needToCheckTop := startY > 2
	needToCheckBottom := startY < len(grid)-3

	if needToCheckTop {
		wg.Add(1)
		go findCharacter(startX, startY, 0, -1, grid, "M", wg, resultChan)
	}

	if needToCheckBottom {
		wg.Add(1)
		go findCharacter(startX, startY, 0, 1, grid, "M", wg, resultChan)
	}

	if needToCheckLeft {
		wg.Add(1)
		go findCharacter(startX, startY, -1, 0, grid, "M", wg, resultChan)

		if needToCheckTop {
			wg.Add(1)
			go findCharacter(startX, startY, -1, -1, grid, "M", wg, resultChan)
		}

		if needToCheckBottom {
			wg.Add(1)
			go findCharacter(startX, startY, -1, 1, grid, "M", wg, resultChan)
		}
	}

	if needToCheckRight {
		wg.Add(1)
		go findCharacter(startX, startY, 1, 0, grid, "M", wg, resultChan)

		if needToCheckTop {
			wg.Add(1)
			go findCharacter(startX, startY, 1, -1, grid, "M", wg, resultChan)
		}

		if needToCheckBottom {
			wg.Add(1)
			go findCharacter(startX, startY, 1, 1, grid, "M", wg, resultChan)
		}
	}
}

func findCrossmas(startX int, startY int, grid [][]string, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()

	canCheckLeft := startX > 0
	canCheckRight := startX < len(grid[0])-1
	canCheckTop := startY > 0
	canCheckBottom := startY < len(grid)-1

	if !(canCheckLeft && canCheckRight && canCheckTop && canCheckBottom) {
		return
	}

	if grid[startY-1][startX-1] == "A" ||
		grid[startY-1][startX+1] == "A" ||
		grid[startY-1][startX-1] == "X" ||
		grid[startY-1][startX+1] == "X" {
		return
	}

	if grid[startY-1][startX-1] == "M" {
		if grid[startY+1][startX+1] != "S" {
			return
		}
	}

	if grid[startY-1][startX-1] == "S" {
		if grid[startY+1][startX+1] != "M" {
			return
		}
	}

	if grid[startY-1][startX+1] == "M" {
		if grid[startY+1][startX-1] != "S" {
			return
		}
	}

	if grid[startY-1][startX+1] == "S" {
		if grid[startY+1][startX-1] != "M" {
			return
		}
	}

	resultChan <- 1
}

func findCharacter(
	startX int,
	startY int,
	directionX int,
	directionY int,
	grid [][]string,
	character string,
	wg *sync.WaitGroup,
	resultChan chan int,
) {
	defer wg.Done()

	if grid[startY+directionY][startX+directionX] == character {
		switch character {
		case "M":
			wg.Add(1)
			go findCharacter(
				startX+directionX,
				startY+directionY,
				directionX,
				directionY,
				grid,
				"A",
				wg,
				resultChan,
			)
		case "A":
			wg.Add(1)
			go findCharacter(
				startX+directionX,
				startY+directionY,
				directionX,
				directionY,
				grid,
				"S",
				wg,
				resultChan,
			)
		case "S":
			resultChan <- 1
		}
	}
}
