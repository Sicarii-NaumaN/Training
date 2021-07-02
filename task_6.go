package main

import (
	"fmt"
	"sync"
	"time"

	"gopkg.in/tomb.v2"
)

func main() {
	quit := make(chan struct{})
	// 1) Способ с помощью управляющего канала
	go func() {
		for {
			select {
			case <- quit:
				return
			default:
				// Что-то делаем
			}
		}
	}()

	//  Что-то делаем
	quit <- struct{}{}

	// 2) С помощью системных вызовов и оповещения об этом всех горутин с помощью контекста(см. 4 задачу)

	// 3) С помощью закрытия канала
	wg := sync.WaitGroup{}
	wg.Add(1)

	ch := make(chan int)
	go func() {
		for {
			smth, ok := <- ch
			if !ok {
				wg.Done()
				return
			}
			fmt.Println(smth)
		}
	}()
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	wg.Wait()

	// 4) С помощью пакета tomb
	proc := tomb.Tomb{}
	proc.Go(func() error{
		for {
			select {
			case <-proc.Dying():
				return nil
			default:
				time.Sleep(300 * time.Millisecond)
				fmt.Println("Я еще жив")
			}
		}
	})

	time.Sleep(1 * time.Second)
	proc.Kill(fmt.Errorf("Потрачено"))
	err := proc.Wait()
	fmt.Println(err)
}
