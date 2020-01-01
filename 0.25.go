package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.println(strings.ContainsAny("Hello World","hi"))
	fmt.println(strings.ContainsAny("Hello World","Hi"))
}