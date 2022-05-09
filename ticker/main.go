package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	for ; true; <-ticker.C {
		fmt.Println("Tick")
	}
}
