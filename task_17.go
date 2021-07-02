package main

import (
	"fmt"
)

func CheckType(obj []interface{}) {		// Функция, проверяющая тип переданных объектов в массиве
	for i := range obj {
		switch obj[i].(type) {
		case int:
			fmt.Println("INT")
		case chan interface{}:
			fmt.Println("CHANNEL")
		case bool:
			fmt.Println("BOOL")
		case string:
			fmt.Println("STRING")
		}
	}
}

func main() {
	i := 1
	b := true
	s1 := ""
	s2 := "1233"
	c  := make(chan interface{}, 1)

	slice := make([]interface{}, 0) // Создаем массив различных из элементов различных типов

	slice = append(slice, i)
	slice = append(slice, b)
	slice = append(slice, s1)
	slice = append(slice, s2)
	slice = append(slice, c)

	CheckType(slice)

}