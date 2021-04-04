package main

import "fmt"

func main (){
	p1 := struct {
		name string
		isMoreThan18 bool
		job string
	}{
		name: "barly",
		isMoreThan18: true,
		job: "web developer",
	}
	fmt.Println(p1)
}
