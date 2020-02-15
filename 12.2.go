package main

import (
	"os"
)

func main()  {
	file, err : os.Create("myFile.txt")
	if err ! = nil {
		return
	}
	defer file.Close()

	fil.WriteString("Hello \n")
	file.WriteString("i am my File.txt")
}