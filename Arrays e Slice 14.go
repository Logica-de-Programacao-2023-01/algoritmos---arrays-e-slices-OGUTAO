package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var i, j int
	fmt.Print("digite o primeiro índice:")
	fmt.Scan(&i)
	fmt.Print("digite o segundo índice:")
	fmt.Scan(&i)

	aux := slice[i]
	slice[i] = slice[j]
	slice[j] = aux

	fmt.Println("slice resultante:", slice)
}
