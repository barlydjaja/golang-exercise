package main

import "fmt"

func main (){
	x := make([]string,50,50)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
