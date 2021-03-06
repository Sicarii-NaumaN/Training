package main

import "fmt"

func main() {
	// Можно сделать еще через удаление всех повторящихся элементов в каждом из массивов,
	// если попросят, напишу.
	// Простой способ через key-value структуру(много вариантов решения)
	// Чтоб не ломать глаза 1, 2, 3 - пересечение
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9 ,9}
	arr2 := []int{1, 10, 34, 3, 16, 30, 2, 34 ,34 ,34}
	tab := make(map[int]*int)

	for i := range arr1 {
		if tab[arr1[i]] == nil {
			counter := 1
			tab[arr1[i]] = &(counter)
		}
	}

	for i := range arr2 {
		if tab[arr2[i]] != nil {
			// Если есть пересечение - инкрементируем
			*tab[arr2[i]]++
		}
	}

	// Если насчитали больше 1 попадания, то записываем в пересечения
	for key, value := range tab {
		if *value > 1 {
			fmt.Printf("%d ", key)
		}
	}
}
