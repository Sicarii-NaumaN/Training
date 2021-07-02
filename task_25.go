package main

import (
	"fmt"
	"sync"
)

// Честно говоря, не очень понял задание
type Counter struct {
	count int
}

// Способ с мьютексами
func main() {
	counter := Counter{}
	n := 10
	m := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			m.Lock()
			counter.count++
			fmt.Println(counter.count)
			m.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()

}

// Способ с каналами
func main2() {
	counter := Counter{}
	n := 10

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
	go func() {
		for range ch {
			counter.count++
			fmt.Println(counter.count)
		}
		chanWg.Done()
	}()
	chanWg.Wait()
}
