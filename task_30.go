package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 5 // удалим, для наглядности, элемент из середины

	// Решение 1 (С сохранением порядка в срезе)
	slice = append(slice[0:i], slice[i+1:]...)

	// Решение 2 (Быстрее, т.к. не нужно копировать массив)
	slice[i] = slice[len(slice)-1]
	slice[len(slice)-1] = 0
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
}
