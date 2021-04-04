package main

import "fmt"

func main(){
	nums := [5]int{}
	nums[0]=1
	nums[1]=2
	nums[2]=3
	nums[3]=4
	nums[4]=5
	fmt.Println(nums)
	fmt.Println(len(nums))
	for i,v := range nums {
		fmt.Println(i,v)
	}
	fmt.Printf("%T", nums)
}
