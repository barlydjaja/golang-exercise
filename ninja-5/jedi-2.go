package main

import "fmt"

type person struct {
	first string
	last string
	favourite []string
}

func main (){
	p1 := person{
		first: "barly",
		last: "djaja",
		favourite: []string{
			"ice cream",
			"chocolate",
			"vanilla",
		},
	}

	p2 := person{
		first: "kara",
		last: "mardika",
		favourite: []string{
			"ice cream",
			"aice",
			"donut",
		},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for key, value := range m{
		fmt.Println(key,value)
	}
}

