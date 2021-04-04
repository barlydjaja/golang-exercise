package main

import "fmt"

// use var to declare variables. package level scope. Don't assign value to variables
// x type int, y type string, z type bool

var x int
var y string
var z bool

func main(){
	fmt.Printf("%v\t",x)
	fmt.Printf("%v\t",y)
	fmt.Printf("%v\t",z)
}
