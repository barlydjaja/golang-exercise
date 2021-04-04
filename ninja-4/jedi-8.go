package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":   []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneyed_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":        []string{"Being Evil", "ice Cream", "Sunsets"},
	}
	for key, value := range m {
		fmt.Println("for name", key)
			for _,description := range value{
				fmt.Println(description)
			}
	}
	delete(m,"bond_james")
	fmt.Println(m)
}
