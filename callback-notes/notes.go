package main

import "fmt"

func sum(xi ...int) int {
	total := 0
	for _, number := range xi {
		total += number
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, number := range vi {
		if number%2 == 0 {
			yi = append(yi, number)
		}
	}
	return f(yi...)
}

func odd(f func(xi ...int) int, v1 ...int) int {
	var yi []int
	for _, number := range v1 {
		if number %2 != 0{
			yi = append(yi, number)
		}
	}
	return f(yi...)
}

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	allSum := sum(ii...)
	fmt.Println(allSum)
	evenSum := even(sum, ii...)
	fmt.Println(evenSum)
	oddSum := odd(sum, ii...)
	fmt.Println(oddSum)
}
