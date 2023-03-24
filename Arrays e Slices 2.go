package main

import "fmt"

func main() {

	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println(numeros)
	numeros = append(numeros[:3], numeros[4:]...)
	fmt.Println(numeros)
}
