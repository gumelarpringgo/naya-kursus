package main

import "fmt"

func IsDivisible(n, x, y int) bool {
	// Create a function that checks if a number n is divisible by two numbers x AND y. All inputs are positive, non-zero numbers.

	// EXAMPLE
	// 1) n =   3, x = 1, y = 3 =>  true because   3 is divisible by 1 and 3
	// 2) n =  12, x = 2, y = 6 =>  true because  12 is divisible by 2 and 6
	// 3) n = 100, x = 5, y = 3 => false because 100 is not divisible by 3
	// 4) n =  12, x = 7, y = 5 => false because  12 is neither divisible by 7 nor 5
	if (n%x == 0) && (n%y == 0) {
		return true
	}
	return false
}

func main() {
	fmt.Println(IsDivisible(3, 3, 4))
	fmt.Println(IsDivisible(12, 3, 4))
	fmt.Println(IsDivisible(8, 3, 4))
	fmt.Println(IsDivisible(48, 3, 4))
	fmt.Println(IsDivisible(100, 5, 10))
	fmt.Println(IsDivisible(100, 5, 3))
	fmt.Println(IsDivisible(4, 4, 2))
	fmt.Println(IsDivisible(5, 2, 3))
	fmt.Println(IsDivisible(17, 17, 17))
	fmt.Println(IsDivisible(17, 1, 17))
}
