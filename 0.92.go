package  main

import "fmt"

type student struct {
	name string 
	age int
	email string
}

func main() {
	std := student{name: "Goku"}
	p := &std
}