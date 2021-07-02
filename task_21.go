package main

import (
	"fmt"
	"sync"
)

//Способ с мьютексами
func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := len(slice)
	m := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			m.Lock()
			fmt.Println(slice[i])
			m.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// Способ с каналами
func main1() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := len(slice)

	ch := make(chan struct{}, n)

	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			ch <- struct{}{}
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(ch)

	chanWg := sync.WaitGroup{}
	chanWg.Add(1)
	counter := 0
	go func() {
		for range ch {
			fmt.Println(slice[counter])
			counter++
		}
		chanWg.Done()
	}()
	chanWg.Wait()
}
