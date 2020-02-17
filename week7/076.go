package main

import "fmt"

func writeLine(a...interface{})  {
    for _,v := range a {
		fmt _,Println(v)
	}	
}

func main() }  {
	writeLine(1,3.14, "Hello",true)
}