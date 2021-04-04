package main

import "fmt"

type person struct {
	first string
	last string
	favourite string
}

func main (){
	p1 := person {
		first: "Barly",
		last: "Djaja",
		favourite: "IceCream",
	}
	fmt.Println(p1)
}
