package main

import "fmt"

func main() {

	array := [5]int{10, 20, 30, 40, 50}

	for _, v := range array {
		fmt.Println(v)
	}
	fmt.Printf("%T", array)
}
