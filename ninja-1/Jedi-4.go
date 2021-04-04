package main

import "fmt"

type hotDog int
var x hotDog
func main (){
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x = 42
	fmt.Println(x)

	y:=int(x)
	fmt.Println(y)
	fmt.Printf("%T",y)
}