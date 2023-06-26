package main

import (
	"fmt"
	"time"
)

func main() {
	x := 1995
	for x <= time.Now().Year() {
		fmt.Println(x)
		x++
	}
}
