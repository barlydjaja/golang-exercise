package main

import "fmt"

func foo(xi ...int)int{
	total := 0
	for _, number := range xi{
		total += number
	}
	return total
}

func bar (x []int) int {
	total := 0
	for _, number := range x{
		total += number
	}
	return total
}

func main (){
	defer fmt.Println("execution done")
	xi := []int{1,2,3,4,5,6,7,8,9,}
	f := foo(xi...)
	b := bar(xi)
	fmt.Println("sum of foo",f)
	fmt.Println("sum of bar", b)
}