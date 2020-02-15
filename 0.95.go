package main

import "fmt"

type I interface{}

fucn main() {
	var i I
	i = "Hollo"
	s, ok := i.(string)
	fmt.Pringln(s, ok)
}