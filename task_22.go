package main

import "fmt"

// Можно и собственный стек впилить, а не колхозный (https://github.com/Sicarii-NaumaN/GOLANG-web-services/blob/task1/calculator/calc/utils.go)
// Нерекурсивный метод
func QuicksortNoRecurse(arr []int, first, last int) {
	stack := make([]int, last-first+1)
	top := 0
	stack[top] = first
	top++
	stack[top] = last - 1

	for top >= 0 {
		last = stack[top]
		top--
		first = stack[top]
		top--

		pivot := partitionNoRecurse(arr, first, last)

		if pivot-1 > first {
			top++
			stack[top] = first
			top++
			stack[top] = pivot - 1
		}

		if pivot+1 < last {
			top++
			stack[top] = pivot + 1
			top++
			stack[top] = last
		}
	}
}

func partitionNoRecurse(arr []int, first, last int) int {
	pivotVal := arr[last]
	i := first - 1
	for j := first; j < last; j++ {
		if arr[j] <= pivotVal {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[last] = arr[last], arr[i+1]
	return i + 1
}

// Классический алгоритм
func Quicksort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	split := partition(arr)

	Quicksort(arr[:split])
	Quicksort(arr[split:])
}

func partition(ar []int) int {
	pivot := ar[len(ar)/2]
	left := 0
	right := len(ar) - 1

	for {
		for ; ar[left] < pivot; left++ {
		}

		for ; ar[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		ar[right], ar[left] = ar[left], ar[right]
	}
}


func main() {
	slice1 := []int{10, 34, 54, 2, 52, -1, 0, 32, 9, 7, 380, 1}
	QuicksortNoRecurse(slice1, 0, len(slice1))
	fmt.Println(slice1)

	slice2 := []int{10, 34, 54, 2, 52, -1, 0, 32, 9, 7, 380, 1}
	Quicksort(slice2)
	fmt.Println(slice2)

	one := 0
	fmt.Println(&one)
	one, two := 1,2
	fmt.Println(&one)
	_ = two
	//fmt.Println(one, two)
	type kek float64
	var g kek = 6.34
	fmt.Printf("%-v", g)
}
