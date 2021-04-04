package main

import "fmt"

func foo() func() int{
	return func() int{
		return 3
	}
}

func main (){
	f := foo()
	fmt.Println(f())
}
