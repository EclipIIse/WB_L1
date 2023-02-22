/*
Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
с инкапсулированными параметрами x,y и конструктором.
*/
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: x,
	}
}

// AB = √(xb - xa)2 + (yb - ya)2
func distanceBetweenPoints(p1, p2 *Point) float64 {
	res := math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
	return res
}

func main() {
	p1 := NewPoint(2, 3)
	p2 := NewPoint(5, 8)

	fmt.Println(distanceBetweenPoints(p1, p2))
}
