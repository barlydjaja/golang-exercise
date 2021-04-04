package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func changeMe(p *person){
	p.last = "joshua"
	fmt.Println(*p)
	(*p).last = "back to djaja"
}

func main(){
	p1 := person{
		first: "barly",
		last: "djaja",
		age: 23,
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
