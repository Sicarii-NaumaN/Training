package main

import (
	"fmt"
	"sync"
)

// 1) Глобальная переменная, которая может являться проблемой при в условиях конкуренции
// (???)* 2) При передачи в функцию элементов меньше 100 (v[:100]) возникнет паника out of range
func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {

		v += "A"
	}
	return v
}

var (
	justString string
	mutex sync.Mutex     // Возможное решение - оборачивание в мьютекс и/или инициализирование
)						 // переменной в скоупе main функции и передача ее явно по указателю


func someFunc() {
	v := createHugeString(1 << 10)

	mutex.Lock()
	justString = v[:100]
	mutex.Unlock()

	fmt.Println(justString)
}

func main() {
	someFunc()
}
