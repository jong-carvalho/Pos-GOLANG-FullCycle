package main

import "fmt"

// Função genérica que recebe um slice de qualquer tipo e imprime os elementos
func printSlice[T any](slice []T) {
	for _, element := range slice {
		fmt.Println(element)
	}
}

func main() {
	intSlice := []int{1, 2, 3, 4}
	stringSlice := []string{"a", "b", "c", "d"}
	boolSlice := []bool{true, false, true}

	fmt.Println("Slice de inteiros:")
	printSlice(intSlice)

	fmt.Println("\nSlice de strings:")
	printSlice(stringSlice)

	fmt.Println("\nSlice de booleanos:")
	printSlice(boolSlice)
}
