package main

import "fmt"

func average ( xs []float64) float64{
	total := 0.0
	for _,v := range xs{
		total += v
	}
	return total / float64(len(xs))
}

func f2() (r int) {
	r = 1
	return
}

func add (args ...int) int{
	total:=0
	for _,v := range args{
		total += v
	}
	return total
}

func main (){
	xs := []float64{
		98,
		93,
		77,
		82,
		83,
	}
	xd := []int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(average(xs))
	fmt.Println(f2())
	fmt.Println(add(xd...))
}