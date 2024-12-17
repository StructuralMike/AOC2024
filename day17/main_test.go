package main

import "testing"

func TestSolvePart1(t *testing.T) {
	sampleDataFile := "sample_input.txt"
	expected := string("4,6,3,5,6,3,5,2,1,0")
	result := solvePart1(sampleDataFile)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
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
