package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 1000)
	go func() {
		for t := range ticker.C {
			fmt.Println("tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 4000)
	ticker.Stop()
	fmt.Println("ticker stopped")
}
