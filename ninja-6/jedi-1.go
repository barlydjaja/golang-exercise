package main

import "fmt"

func foo () int {
	return 3
}

func bar () (int, string){
	return 4, "this is number four"
}

func main (){
	f := foo()
	bnum, bword := bar()
	fmt.Println(f, bnum, bword)
}
