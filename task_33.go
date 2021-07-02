package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Не знаю как легче это реализовать, буду рад если поможете

// Генерируем рандомный seed для недетерминированного рандома
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Функция достает из первого канала число, проверяет его на четность, и записывает четные во второй
func checkOdd(sender <-chan int, receiver chan<- int) {
	go func() {
		for {
			select {
			case d := <-sender: // Получение из первого канала данных
				if d%2 == 0 { // Проверка на четность и запись
					receiver <- d
				}
			}
		}
	}()
}

func main() {
	chanFirst, chanSecond := make(chan int), make(chan int)
	stop := make(chan struct{}) // Канал, информурующий о том, что данные в первом канале закончились
	go func() {
		for i := 0; i < 10; i++ { // Взял для примера 10 чисел
			chanFirst <- rand.Intn(100)
		}
		stop <- struct{}{}
	}()

	checkOdd(chanFirst, chanSecond)

	defer close(chanFirst)
	fmt.Println("Result:")
	for {
		select {
		case out := <-chanSecond:
			fmt.Println(out) // Выводим результат
		case <-stop:
			close(chanSecond)
			return
		}
	}
}
