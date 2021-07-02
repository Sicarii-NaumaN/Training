package main

import "fmt"

// Выведутся 2 единицы, так как в GO все передается по значению, даже указатели, то есть в
// функции update просто создается копия значения указателя и ее мы локально в скоупе функции
// переписваиваем новым указателем, однако в main p имеет тот же адрес, мы може  таким образом
// переписать только значение p:  *p = b
func update(p *int) {
	fmt.Println("Ptr update: ", p)
	b := 2
	p = &b
	fmt.Println("Ptr after update: ", p)
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
}
