package main

import "fmt"

func main() {
	inc := incrementer()

	fmt.Println(inc()) // Output: 1
	fmt.Println(inc()) // Output: 2
	fmt.Println(inc()) // Output: 3

	anotherInc := incrementer()
	fmt.Println(anotherInc()) // Output: 1
	fmt.Println(anotherInc()) // Output: 2

	fmt.Println(inc()) // Output: 4
}

func incrementer() func() int {
	count := 0 // Variável para armazenar o estado interno
	return func() int {
		count++ // Incrementa o contador
		return count
	}
}

//Implemente uma função que retorne outra função. A função retornada deve incrementar um número inteiro toda vez que for chamada.
