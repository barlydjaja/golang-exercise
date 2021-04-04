package main

import "fmt"

func sum(xi ...int) int{
	total := 0
	for _,number := range xi{
		total += number
	}
	return total
}

func even(f func(xi ...int)int, xi ...int) int{
	var yi []int
	for _,number := range xi{
		if number % 2 == 0{
			yi = append(yi,number)
		}
	}
	return f(yi...)
}

func main (){
	xi := []int{1,2,3,4,5,6,7,8,9,}
	sumTotal := sum(xi...)
	evenTotal := even(sum,xi...)
	fmt.Println(sumTotal)
	fmt.Println(evenTotal)
}
