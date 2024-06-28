package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := 3

	if a > b {
		fmt.Println(a)
	} else if b > c {
		fmt.Println(b)
	} else {
		fmt.Println(c)
	}

	switch a {
	case 10:
		fmt.Println("a")
	case 20:
		fmt.Println("b")
	default:
		fmt.Println("d")
	}
}
