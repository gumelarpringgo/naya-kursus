package main

import "fmt"

func main() {
	x := 11
	y := 88

	fmt.Println("Variable x =", x)
	fmt.Println("Variable y =", y)

	fmt.Println("SWAP")

	x, y = y, x
	fmt.Println("Variable x =", x)
	fmt.Println("Variable y =", y)
}
