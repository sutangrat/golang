package main

import (
	"fmt"
	"sync"
	"time"
)	

func say(txt string, sleep time.Duration,wg *sync.WaiGroup)  {
	defer wg. Done()
	fmt.Println(txt)
	time.Sleep(time.Millisecond * sleep)
}

func main() {
	var wg sync.WaiGroup
	wg.Add (2)
	
}