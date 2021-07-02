package main

import "fmt"

// Конвейер, записываем во 2 канал данные(x*2)
func conveyor(ch1 <-chan int) <-chan int{
	ch2 := make(chan int)
	go func() {
		defer close(ch2)
		for {
			select {
			case d := <-ch1:
				ch2 <- 2 * d
			}
		}
	}()
	return ch2
}

func main() {
	// Наш произвольный массив
	arr := []int{2, 4, 6, 8, 10}
	// Первый канал
	ch1 := make(chan int)
	// Второй канал
	ch2 := conveyor(ch1)

	// Заполняем первый канал
	go func() {
		for i := range arr {
			ch1 <- arr[i]
		}
		close(ch1)
	}()

	// Выводим результат второго канала
	for range arr {
		fmt.Println(<-ch2)
	}
}
