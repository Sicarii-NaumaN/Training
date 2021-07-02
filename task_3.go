package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	sum := 0
	//m := new(sync.Mutex)
	arr := []int{2, 4, 6, 8, 10}
	for i := range arr {
		wg.Add(1)
		go func(number int, sum *int) {
			defer wg.Done()
			//m.Lock()
			*sum += number*number
			//m.Unlock()
		}(arr[i], &sum)
	}

	wg.Wait()
	fmt.Println(sum)
}