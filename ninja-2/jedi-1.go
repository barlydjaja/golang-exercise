package main

import "fmt"

//typed constant
const age int = 21

// untyped constant
const randomNum = 21

func main(){
	num := 2

	fmt.Printf("%d \t %b \t %#x \n", num, num, num)
	fmt.Println(age, randomNum)
}
