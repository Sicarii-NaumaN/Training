package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	d := time.NewTimer(duration)
	select { // Будет висеть, т.к. ждет выполнения какого-либо кейса,
	case <-d.C: // У нас нет default кейса, поэтому цикл не нужен
		 // Отдельная горутина также не нужна т.к нам нужна блокировка работы
	}
}

func Sleep1(duration time.Duration) { // Не для отрицательных значений
	select {
	case <-time.After(duration):
	}
}

func Sleep2(duration time.Duration) {// Не для отрицательных значений
	d := time.NewTicker(duration)
	select {
	case <-d.C:
		//d.Stop()
	}
}


func main() {
	start := time.Now()
	Sleep(3 * time.Second)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
