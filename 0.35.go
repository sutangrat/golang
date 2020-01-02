package main

import "fmt"

func main() {
	fmt.Print("input : ")
	var name string
	var height float32
	var Weigth float32
	n, err := fmt.Scan(&name, &age, &Weigth, &height)
	fmt.Println(name, age, weigth, height)
	fmt.Println('number of argument ', n)
}