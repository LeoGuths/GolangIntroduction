package Ninja2

import "fmt"

func main() {
	x := 5
	y := 6

	a := x == y
	b := x != y
	c := x <= y
	d := x < y
	e := x >= y
	f := x > y

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
}
