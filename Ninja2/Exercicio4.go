package Ninja2

import "fmt"

func main() {
	x := 10
	fmt.Printf("%d\n%b\n%#x\n\n", x, x, x)

	y := x << 1
	fmt.Printf("%d\n%b\n%#x", y, y, y)
}
