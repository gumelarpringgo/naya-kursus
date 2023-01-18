package init

import "fmt"

var y int = 10
var x string = "gumelar"
var z bool = true

func init() {
	fmt.Println(y)
	fmt.Println(x)
	fmt.Println(z)

}

// package command-line-arguments is not a main package
