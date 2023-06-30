package main

import "fmt"

func main() {
	elementos := [4]float64{1, 1.2, 1.5, 1.7}
	var conta float64
	conta = 1
	for _, x := range elementos {
		conta = conta * x
	}
	fmt.Println("O produto é", conta)
}
