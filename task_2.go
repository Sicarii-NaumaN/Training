package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	arr := []int{2, 4, 6, 8, 10}
	for i := range arr {
		wg.Add(1)
		go func(number int) {
			defer wg.Done()
			fmt.Println(number*number)
		}(arr[i])
	}

	 wg.Wait()
}
