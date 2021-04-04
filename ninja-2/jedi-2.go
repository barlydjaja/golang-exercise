package main

import "fmt"

func main(){
	num := 23
	fmt.Printf("decimal: %d\nbinary:%b\nhex:%#x\n", num, num ,num)
	shiftedNum := num<<1
	fmt.Printf("decimal: %d\nbinary:%b\nhex:%#x\n", shiftedNum, shiftedNum ,shiftedNum)
}
