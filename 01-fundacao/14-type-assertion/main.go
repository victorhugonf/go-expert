package main

import "fmt"

func main() {
	var v interface{} = "VH"
	println(v)
	println(v.(string))

	res, ok := v.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)

	res2 := v.(int)
	fmt.Printf("O valor de res é %v",res2)
}
