package main

import (
	"testing"
)

type Test struct {
	q int
	a int
}

func test(f Problem, tests []Test, t *testing.T) {
	for i := range tests {
		answer := f(tests[i].q)

		if answer != tests[i].a {
			t.Errorf("Expected %d, got %d\n", tests[i].a, answer)
		}
	}
}

func TestMultiplesOf3And5(t *testing.T) {
	tests := []Test{
		{1000, 233168},
		{500, 57918},
		{250, 14543},
		{100, 2318},
		{50, 543},
	}

	f := MultiplesOf3And5

	test(f, tests, t)

}

func TestEvenFibbonacciNumbers(t *testing.T) {
	tests := []Test{
		{4000000, 4613732},
	}

	f := EvenFibbonacciNumbers

	test(f, tests, t)
}

func TestLargestPrimeFactor(t *testing.T) {
	tests := []Test{
		{10, 5},
		{100, 5},
		{14, 7},
		{18, 3},
		{38, 19},
	}

	f := LargestPrimeFactor

	test(f, tests, t)
}
