package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

type jsonPerson struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	// marshalling
	p1 := person{
		First: "Barly",
		Last:  "Djaja",
		Age:   23,
	}
	p2 := person{
		First: "Kartika",
		Last:  "Mardika",
		Age:   25,
	}
	people := []person{p1, p2}
	marshalBS, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshalBS))

	// unmarshalling
	s := `[{"First":"Barly","Last":"Djaja","Age":23},{"First":"Kartika","Last":"Mardika","Age":25}]`
	bs := []byte(s)
	fmt.Printf("%T\n", bs)

	var jsonPeople []jsonPerson

	err = json.Unmarshal(bs, &jsonPeople)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(jsonPeople)

}
