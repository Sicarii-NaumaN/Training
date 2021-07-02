package main

import (
	"fmt"
	"sync"
)
// Выведется все, кроме постледнего fmt.Println("exit").
// Так как мы передаем копию объекта WaitGroup, то возникает deadlock в wg.Wait, потому что
// наша оснавная горутина не может дождаться окончания работы всех горутин и счетчик горутин не уменьшается
// Для исправления ошибки следует передавать WaitGroup поссыл ке в другие горутины, либо
// не передавать ее вовсе, так как они в одной области видимости
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)

		// TODO Решение 1
		//go func(wg *sync.WaitGroup, i int) {
		//	fmt.Println(i)
		//	wg.Done()
		//}(&wg, i)

		// TODO Решение 2
		//go func(i int) {
		//	fmt.Println(i)
		//	wg.Done()
		//}(i)

	}
	wg.Wait()
	fmt.Println("exit")
}

