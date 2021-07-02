package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9 ,9}
	arr2 := []int{1, 10, 1, 3, 12, 11, 2, 10 , 10, 4}
	merge := make(map[int]*struct{})
	isElem := struct{}{}

	for i := range arr1 {
		if merge[arr1[i]] == nil {
			merge[arr1[i]] = &isElem
		}
	}

	for i := range arr2 {
		if merge[arr2[i]] == nil {
			merge[arr2[i]] = &isElem
		}
	}

	for key, _ := range merge {
		fmt.Printf("%d ", key)
	}
}
