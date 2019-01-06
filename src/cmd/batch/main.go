package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("batch")
		time.Sleep(5 * time.Second)
	}
}
