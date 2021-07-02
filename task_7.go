package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	elemCount := 10
	m := new(sync.Mutex)

	testMap := make(map[int]int)
	for i := 0; i < elemCount; i++ {
		wg.Add(1)
		go func(test map[int]int, anyData int) {
			defer wg.Done()

			m.Lock()
			test[anyData] = anyData
			m.Unlock()

		}(testMap, i)
	}

	wg.Wait()
	fmt.Println(testMap)
}