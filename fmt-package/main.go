package main

import (
	"fmt"
)

var a int
var b = "James Bond"
var c bool
var d = true

func main() {
	e := 42
	f := "Shaken not stirred"
	g := `Miss Moneypenny says, "Hellooo, James."`
	h := `Q says,
	"I have some new toys for you, James."`

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(b, "says,", f)
	fmt.Println(g)
	fmt.Println(h)

	s := fmt.Sprint(a, " semething more ", b)
	fmt.Println(s)
}
