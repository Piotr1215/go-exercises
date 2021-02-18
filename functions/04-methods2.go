package functions

import (
	"fmt"
	"math"
)

type square struct {
	length int
	width  int
}

type circle struct {
	radius int
}

type shape interface {
	area() float32
}

func Entry() {

	square := square{
		100,
		200,
	}

	circle := circle{
		100,
	}

	info(square)
	info(circle)
}

func info(s shape) {
	fmt.Println("The area is", s.area())
}

func (c circle) area() float32 {
	return math.Pi * (float32(c.radius) * float32(c.radius))
}

func (s square) area() float32 {
	return float32(s.length) * float32(s.width)
}
