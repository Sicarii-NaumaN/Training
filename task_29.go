package main

import (
	"fmt"
	"log"
	"math/big"

)
// Простенький калькулятор для больших чисел.
func main() {
	var number1, number2 string
	fmt.Println("Enter numbers:")
	_, err := fmt.Scan(&number1, &number2)
	if err != nil {
		log.Fatal(err)
	}

	bigNumber1 := new(big.Float)			// Инициализируем наши числа
	_, err = fmt.Sscan(number1, bigNumber1)
	if err != nil {
		log.Fatal(err)
	}

	bigNumber2 := new(big.Float)
	_, err = fmt.Sscan(number2, bigNumber2)
	if err != nil {
		log.Fatal(err)
	}

	res := new(big.Float)		// Результирующее число
	fmt.Println("Enter action \nmultiply: 1, \ndivide: 2, \nplus: 3, \nminus: 4")
	choice := 0
	_, err = fmt.Scan(&choice)
	if err != nil {
		log.Fatal(err)
	}

	switch choice {          // Выбор действия над числами (не стал выносить в отдельную функцию)
	case 1:
		fmt.Println(res.Mul(bigNumber1, bigNumber2))
	case 2:
		fmt.Printf("%.20f\n", res.Quo(bigNumber1, bigNumber2))
	case 3:
		fmt.Println(res.Add(bigNumber1, bigNumber2))
	case 4:
		fmt.Println(res.Sub(bigNumber1, bigNumber2))
	default:
		fmt.Println("Not supported action passed")
	}
}