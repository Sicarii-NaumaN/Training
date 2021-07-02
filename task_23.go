package main

import "fmt"

func BinarySearch(arr []int, searchElem int) int {
	first, last := 0, len(arr)

	// Чтоб за массив не убежали, запомним изначальный размер
	initLen := len(arr)

	for first < last {
		mid := (first + last) / 2
		if arr[mid] < searchElem {
			first = mid + 1
		} else {
			last = mid
		}
	}

	if first == initLen || arr[first] != searchElem {
		return -1
	}
	return first
}

func main() {
	// Отсортированный массив
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(BinarySearch(s, 2))
}

