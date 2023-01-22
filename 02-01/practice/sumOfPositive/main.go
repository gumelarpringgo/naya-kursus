package main

import "fmt"

func PositiveSum(numbers []int) int {
	var x int
	for _, num := range numbers {
		if num > 0 {
			x += num
		}
	}
	return x
}
func cob() {
	// jika suatu fungsi tidak terdapat nilai, maka dapat berjalan. jika fungsi terdapat nilai, maha harus digunakan
}

func main() {
	numbers := []int{1, 3, 5, -3, 4, -7}
	fmt.Println(PositiveSum(numbers))
}
