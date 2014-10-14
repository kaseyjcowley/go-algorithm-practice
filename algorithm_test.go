package algorithm

import (
	"log"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

type TestPair struct {
	values []int
	answer []int
}

var (
	tests    []TestPair
	numTests int = 5
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < numTests; i++ {
		testArray := make([]int, rand.Intn(SIZE)+1)

		for j := 0; j < len(testArray); j++ {
			testArray[j] = rand.Intn(SIZE)
		}

		sortedArray := make([]int, len(testArray))
		copy(sortedArray, testArray)

		sort.Ints(sortedArray)

		test := TestPair{
			testArray,
			sortedArray,
		}

		tests = append(tests, test)
	}
}

func TestQuickSort(t *testing.T) {
	for _, pair := range tests {
		start := time.Now()
		Quicksort(pair.values)
		duration := time.Since(start)

		log.Printf("Quicksort on %d items took %f seconds", len(pair.values), duration.Seconds())

		if !reflect.DeepEqual(pair.values, pair.answer) {
			t.Errorf("Expected %v, got %v", pair.answer, pair.values)
		}
	}
}
