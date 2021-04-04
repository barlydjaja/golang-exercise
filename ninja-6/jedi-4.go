package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p person) speak (){
	fmt.Println("this person named ", p.first, " can speak, is of age ", p.age)
}

func main (){
	p1 := person {
		first: "barly",
		last: "djaja",
		age: 23,
	}
	p1.speak()
}
