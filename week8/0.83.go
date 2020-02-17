package main

import "fmt"

func handlePanic() {
	r := recover()
	if r == "to munh" {
		fmt.Print("your number our of range")
		main()
	}
}
func main(){
	defer handlePanic ()
	var i int
	fmt.Print("type number :")
	_, e := fmt.Scan(&i)
	if e ! = nit {
		panic("to much")
	}	
	fmt.Println("your number :", i)
}	

