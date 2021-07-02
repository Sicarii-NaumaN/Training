package main

import (
	"fmt"
	"log"
)

// Проверяем бит
func isBit(n uint, pos uint) bool {
	val := n & (1 << pos)
	return val > 0
}

// 0 --> 1
func setBit(n uint, pos uint) uint {
	n |= 1 << pos
	return n
}

// 1 --> 0
func clearBit(n uint, pos uint) uint {
	mask := uint(^(1 << pos))
	n &= mask
	return n
}

func main() {
	var number int64
	var i uint
	fmt.Println("Enter number and bit")
	_, err := fmt.Scan(&number, &i)
	if err != nil {
		log.Fatal(err)
	}

	// Если на указанной позиции есть бит, то зануляем его, иначе ставим единицу
	if isBit(uint(number), i) {
		res := clearBit(uint(number), i)
		fmt.Printf("%b int:%d\n", res, res)
	} else {
		res:= setBit(uint(number), i)
		fmt.Printf("%b int:%d\n", res, res)
	}
}