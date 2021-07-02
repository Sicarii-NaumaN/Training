package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Ну если правильно понял задание ;)

// Наши worker-ры выводят рандомные данные, и слушают сигнал из контекста
func worker(ctx context.Context, wg *sync.WaitGroup, dataCh chan int) {
	defer wg.Done()
	for {
		select {
		case task, ok := <-dataCh:
			if !ok {
				return
			}
			fmt.Println("Received data: ", task)
		// Если пришло оповещение о terminate основной горутины, то все горутины завершают работу
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	workers := 0

	fmt.Println("Enter worker's count")
	_, err := fmt.Scan(&workers)
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(workers)
	// Основной канал с данными
	dataChan := make(chan int)
	for i := 0; i < workers; i++ {
		go worker(ctx, &wg, dataChan)
	}

	termChan := make(chan os.Signal)
	// Подписываемся на события
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	defer wg.Wait()
	defer close(dataChan)

	exit := false
	go func() {
		select {
		// Если пользователь захочет прервать программу
		case sig := <-termChan:
			if sig == syscall.SIGINT || sig == syscall.SIGTERM {
				cancelFunc()
				exit = true

				// Чтоб горутина осознала свой конец, и не писала в закрытый канал
				d := time.Duration(100) * time.Millisecond
				time.Sleep(d)
				close(termChan)
				close(dataChan)

				return
			}
		}
	}()

	// Ну и цикл, пока пользователь его не прибьет
	for exit != true {
		dataChan <- rand.Intn(100)
	}
}
