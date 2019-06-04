package main

import "fmt"

func binarySearch(arr []int, element int) (bool, int) {
	end := len(arr) - 1
	median := end / 2
	for start := 0; start <= end; {
		median = (start + end) / 2
		fmt.Printf("start: %d, median: %d, end: %d \n", start, median, end)

		val := arr[median]
		if element == val {
			return true, median
		} else if element > val {
			start = median + 1
		} else if element < val {
			end = median - 1
		}
	}
	return false, -1
}

func testBinarySearch(element int) {
	var arr []int
	for i := 1; i <= 100; i++ {
		arr = append(arr, i)
	}
	ok, index := binarySearch(arr, element)
	if ok {
		fmt.Printf("Success, the index of the element is %d \n", index)
	} else {
		fmt.Printf("Failed, this element was not found: %d \n", element)
	}
}

func main() {
	testBinarySearch(1)
	testBinarySearch(23)
	testBinarySearch(100)
	testBinarySearch(101)
}
