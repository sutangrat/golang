package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os. Open(".")
	if err != nil {
		return
	}
	defer dri.Close()

	fileInfos, err := dir.Readdir(-1)
	for _, fi := range fileInfos {
		filename := fi.Name()
		fmt.Println(filename)
	}
}