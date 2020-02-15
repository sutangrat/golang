package main

import (
	"fmt"
	"os"
	"path/filpath"
)

func walkFn(path string, info os. FileInfo, err error) error {
	fmt.Println(path)
	return nil
}
func main()  {
	filpath.Walk(".",walkFn)
}