package algorithm

const SIZE = 100

// Quicksort takes an array of integers and sorts them
func Quicksort(array []int) {
	// Never sort an array of 0 or 1 length
	if len(array) < 2 {
		return
	}

	// Select the pivot and determine the left and right pointers
	pivot := array[0]
	left, right := 1, len(array)-1

	// Continue looping until the left pointer hits the right
	for left <= right {
		// If the number is less than the pivot, continue the low pointer
		for left <= right && array[left] <= pivot {
			left++
		}

		// If the number is greater than the pivot, continue the high pointer
		for right >= left && array[right] >= pivot {
			right--
		}

		// Peform the swap
		if left < right {
			array[left], array[right] = array[right], array[left]
		}
	}

	if right > 0 {
		array[0], array[right] = array[right], array[0]
		// Quicksort(array[0:right])
	}

	Quicksort(array[left:])
}
