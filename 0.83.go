package main

import "fmt"

func handlePanic() {
	r := recover()
	if r == "to munh" {
		fmt.Print("your number our of range")
	}
}

