package main
import "fmt"

func main() {
	alphabets := [4]string{"A", "B", "C", "D"}
	fmt.Println(alphabets)

	x := alphabets[0:2]
	fmt.Println(x)

	z := y[0:1]
	fmt.Println(z)

	z[0] = "x"
	
}