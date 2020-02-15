package main

import (
	"fmt"
	"sync"
	"time"
)	

func say(txt string, sleep time.Duration,wg *sync.WaiGroup)  {
	defer wg. Done()
}