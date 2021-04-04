package main

import "fmt"

func incrementer() func() int{
	x := 0
	return func() int {
		x++
		return x
	}
}

func main (){
	f := incrementer()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
