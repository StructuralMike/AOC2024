package main

import "testing"

func TestSolvePart1(t *testing.T) {
	sampleDataFile := "sample_input.txt"
	expected := 7036
	result := solvePart1(sampleDataFile)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
func TestSolvePart1_2(t *testing.T) {
	sampleDataFile := "sample_input2.txt"
	expected := 11048
	result := solvePart1(sampleDataFile)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// func TestSolvePart2(t *testing.T) {
// 	sampleDataFile := "sample_input.txt"
// 	expected := 0
// 	result := solvePart2(sampleDataFile)
// 	if result != expected {
// 		t.Errorf("Expected %d, got %d", expected, result)
// 	}
// }
