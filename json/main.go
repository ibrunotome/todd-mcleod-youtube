package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"Fist"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Wonder","Last":"Woman", "Age": 20}]`
	bs := []byte(s)

	// fmt.Println("%T\n", s)
	// fmt.Println("%T\n", bs)

	people := []person{}

	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nall of the data", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
