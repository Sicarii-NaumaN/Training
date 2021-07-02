package main

import "fmt"

func main() {
	symbolsTab := make(map[rune]bool)
	str := ""
	fmt.Scan(&str)
	symbols := []rune(str)
	for i := range symbols {
		if symbolsTab[symbols[i]] {			// Если символ уже есть в таблице, то завершаем программу
			fmt.Println("Not unique")
			return
		}
		symbolsTab[symbols[i]] = true
	}
	fmt.Println("Unique")
}