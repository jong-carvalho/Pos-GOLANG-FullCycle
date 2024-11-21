package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	runsTroughSlice(s)

}

func runsTroughSlice(s []int) {
	var evenNumbers = []int{}

	for i := 0; i < 10; i++ {
		if s[i]%2 == 0 {
			evenNumbers = append(evenNumbers, s[i])
		}
	}

	fmt.Printf("Tamanho=%d Capacidade=%d %v\n", len(evenNumbers), cap(evenNumbers), evenNumbers)
}

//Crie um slice que contenha uma lista de 10 números.
//Implemente uma função que percorra o slice e retorne um novo slice apenas com números pares.
