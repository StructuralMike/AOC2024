package main

import "testing"

func TestSolvePart1(t *testing.T) {
	sampleDataFile := "sample_input.txt"
	expected := 2
	result := solvePart1(sampleDataFile)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	sampleDataFile := "sample_input.txt"
	expected := 4
	result := solvePart2(sampleDataFile)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
