package main

import "fmt"

func main() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tab := make(map[int][]float64)

	for i:= range temp {
		// Определяем интервал, избавляемся от дробной части
		interval := int(temp[i])/10*10
		tab[interval] = append(tab[interval], temp[i])
	}

	fmt.Println(tab)
}