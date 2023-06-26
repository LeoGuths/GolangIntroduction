package main

import (
	"fmt"
	"time"
)

func main() {
	x := 1995
	for {
		fmt.Println(x)
		if x == time.Now().Year() {
			break
		}
		x++
	}
}
