package main

import (
	"fmt"
	"stings"
)

func main() {
	fmt.Println(stings.Replace("Hello World","1","x",2))
	fmt.Println(stings.Replace("Hello World","1","x",-1))
}