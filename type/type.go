package main

import (
	"fmt"
	"math/cmplx"
)

var c, python, java bool
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println("========================= TYPE ===========================")
	fmt.Println(split(17))
	var i int
	var a, b int = 1, 2
	c, python, java := true, false, "no!"
	fmt.Println(i, c, python, java)
	fmt.Println(a, b)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
