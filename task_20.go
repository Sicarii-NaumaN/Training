package main

import "fmt"

// Исходный слайс не изменится никак, так как анонимная функция соддает абсолютно новый слайс,
// скопированный с переданного(в отличие от обычной функции). Достаточно не передавать в
// анонимную функцию слайс(т.к. одна область видимости)
func main() {
	slice := []string{"a", "a"}
	fmt.Printf("main %p %p\n ", slice, &slice[0])

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Println(slice)
		fmt.Printf("anonymous %p %p\n ", slice, &slice[0])
	}(slice)
	fmt.Println(slice)

}
