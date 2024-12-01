package main

import (
	"os"
	"testing"
)

func TestSolve(t *testing.T) {
	data, err := os.ReadFile("example.txt")
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	answer1, time1, answer2, time2 := solve(string(data))

	expectedPart1 := 11
	expectedPart2 := 31

	if answer1 != expectedPart1 {
		t.Fatalf("Part 1 failed: got %d, expected %d", answer1, expectedPart1)
	}

	t.Logf("Part 1 succeeded: got %d in %dms", answer1, time1)

	if answer2 != expectedPart2 {
		t.Fatalf("Part 2 failed: got %d, expected %d", answer2, expectedPart2)
	}

	t.Logf("Part 2 succeeded: got %d in %dms", answer2, time2)
}
