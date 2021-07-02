package main

import "fmt"

func main() {
	n := 0
	if true { // Выведется 0, т.к. мы создаем новую переменную оператором :=
		n := 1 // в теле условия
		n++
	}
	fmt.Println(n)
}
