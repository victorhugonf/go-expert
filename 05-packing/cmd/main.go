package main

import (
	"fmt"

	"github.com/victorhugonf/go-expert/05-packing/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())

	m2 := math.NewMathPrivateFields(3, 4)
	fmt.Println(m2.Add())

	m3 := math.NewMathPrivate(5, 6)
	fmt.Println(m3.Add())
	fmt.Println(m3)
}
