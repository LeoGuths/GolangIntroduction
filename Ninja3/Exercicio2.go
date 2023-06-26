package main

import (
	"fmt"
)

func main() {
	for x := 65; x <= 90; x++ {
		fmt.Println(x)
		for y := 0; y <= 2; y++ {
			fmt.Printf("\t%#U \t \n", x)
		}
		fmt.Printf("\n")
	}
}
