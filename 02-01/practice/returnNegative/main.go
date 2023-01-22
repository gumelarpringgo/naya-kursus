package main

import "fmt"

func MakeNegative(x int) int {
	if x > 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(MakeNegative(5))
	fmt.Println(MakeNegative(-3))
	fmt.Println(MakeNegative(5555))
	fmt.Println(MakeNegative(0))

}
