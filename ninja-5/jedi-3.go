package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main (){
	vehicle1 := truck{
		vehicle:vehicle{
			doors: 4,
			color: "yellow",
		},
		fourWheel: true,
	}
	vehicle2 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		luxury: false,
	}

	fmt.Println(vehicle1)
	fmt.Println(vehicle2)
	fmt.Println(vehicle1.doors)
	fmt.Println(vehicle2.doors)
}
