package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func factorial(input int) int {
	if input == 1 || input == 0 {
		return input
	}
	return input * factorial(input-1)
}

func main() {
	for {
		fmt.Print("Enter your factorial number: ")

		read := bufio.NewReader(os.Stdin)

		input, _ := read.ReadString('\n')
		input = input[:len(input)-1]
		inputFactorial, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Enter number only")
			continue
		}

		value := factorial(inputFactorial)
		fmt.Printf("Factorial %d = %d\n", inputFactorial, value)
		break
	}

}
