package main

import (
	"testing"
)

func TestMultiplesOf3And5(t *testing.T) {
	correct := 233168
	answer := MultiplesOf3And5()

	if correct != answer {
		t.Errorf("Expected %d, got %d\n", correct, answer)
	}
}

func TestLargestPrimeFactor(t *testing.T) {
	correct := 6857
	answer := LargestPrimeFactor()

	if correct != answer {
		t.Errorf("Expected %d, got %d\n", correct, answer)
	}
}
