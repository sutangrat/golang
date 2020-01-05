package main

import "fmt"

func main() {
	elements := make(map[sting]sting)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"

	h := elements["H"]
	fmt.Println(h)

	n, found := elements["N"]
	fmt.Println(n)
}