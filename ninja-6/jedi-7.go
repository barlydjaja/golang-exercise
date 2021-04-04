package main

import "fmt"

func main() {

	x := 0
	f := func() {
		x++
		fmt.Println(x)
	}
	f()
	f()
	f()
	f()
}
