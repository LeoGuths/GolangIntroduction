﻿package main

import "fmt"

const (
	_ = 2023 + iota
	a
	b
	c
	d
)

func main() {
	fmt.Printf("%v\n%v\n%v\n%v\n", a, b, c, d)
}
