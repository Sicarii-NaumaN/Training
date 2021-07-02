package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	dataCh := make(chan int)
	fmt.Println("Enter time in seconds")
	n := time.Duration(0)
	fmt.Scan(&n)
	timer := time.After(n * time.Second)  // Можно и с тиками сделать и newTimer(нет garbage коллектора)
	stop := false						// Флаг остановки

	// Вариант с 3 горутинами
	go func() {
		select {
		case <-timer:
			stop = true
			return
		}
	}()

	go func() {
		for !stop {
			dataCh <- rand.Intn(100)
		}
		close(dataCh)
	}()

	// Вариант с двумя горутинами(бесконечный цикл внутри)

	//go func() {
	//	for {
	//		select {
	//		case <-timer:
	//			stop = true
	//			close(dataCh)
	//			return
	//		default:
	//			dataCh <- rand.Intn(100)
	//		}
	//	}
	//
	//}()

	for !stop {
		fmt.Println(<-dataCh)
	}
}
