package main

import "fmt"

func sum (numbers...int) int {
	total := 0
	for _, n := range numbers{
		total = total + n
	}
	return total
}

func main() {
	x := (2,4,6,8)
	fmt.println(x)

	y := sum()
	
}