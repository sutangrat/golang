package main

import "fmt"

const greeting = "Hello World"

func main() {
	fmt.Println(greeting)
	greeting = "Hallo welt"
	fmt.Println(greeting)
}