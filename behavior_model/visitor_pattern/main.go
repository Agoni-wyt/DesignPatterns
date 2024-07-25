package main

import (
	"fmt"
	"visitor_pattern/visitor"
)

func main() {
	square := visitor.NewSquare(2)
	circle := visitor.NewCircle(3)
	rectangle := visitor.NewRectangle(2, 3)

	areaCalculator := &visitor.AreaCalculator{}

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &visitor.MiddleCoordinates{}
	square.Accept(middleCoordinates)
	circle.Accept(middleCoordinates)
	rectangle.Accept(middleCoordinates)
}
