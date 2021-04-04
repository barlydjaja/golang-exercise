package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64{
	return math.Pi*c.radius*c.radius
}

func (s square) area() float64{
	return s.length*s.length
}

type shape interface {
	area() float64
}

func info(s shape){
	x := s.area()
	fmt.Println(x)
}

func main (){
	c := circle{
		radius: 5,
	}
	s := square{
		length: 2,
	}
	info(c)
	info(s)
}
