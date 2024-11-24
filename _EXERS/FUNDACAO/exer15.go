package main

import "fmt"

func main() {

	var number int
	fmt.Print("Digite um número: ")
	fmt.Scanln(&number)

	fmt.Print(intNumberPointer(&number))

}

func intNumberPointer(number *int) int {
	return *number * 2
}

//Crie uma função que receba um ponteiro para um número inteiro e atualize o valor para o dobro do original.
