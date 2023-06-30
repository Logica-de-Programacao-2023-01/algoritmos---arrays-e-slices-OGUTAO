package main

import "fmt"

func main() {
	var soma, tamanho, x int
	var slice []int

	fmt.Print("qual o tamanho desse slice")
	fmt.Scan(&tamanho)
	for i := 0; i < tamanho; i++ {

		fmt.Print("digite um número:")
		fmt.Scan(&x)
		slice = append(slice, x)
		soma += x
	}

	fmt.Println("o SLice é", slice)
	fmt.Println("A soma é de", soma)
}
