package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Enter your username: ")

	read := bufio.NewReader(os.Stdin)

	input, _ := read.ReadString('\n')

	fmt.Printf("Helloo %s!\n", strings.TrimSpace(input))
}
