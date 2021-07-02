package main

import (
	"fmt"
	"math"
)

func NewPoint(_x, _y float64) *Point {
	return &Point{x: _x, y: _y}
}

type Point struct {
	x float64
	y float64
}

func GetDistance(p1, p2 *Point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)) // Можно было и с math.Pow
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(6, 7)
	fmt.Println(GetDistance(p1, p2))
}