package main

import "fmt"

type I interface{}

fucn main() {
	var i I
	i = "Hollo"
	s, ok := i.(string)
	fmt.Pringln(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	b := i. (bool)
	fmt.Println(b)
}