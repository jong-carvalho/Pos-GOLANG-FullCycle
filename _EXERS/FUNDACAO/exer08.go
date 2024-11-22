package main

import "fmt"

func main() {

	var num1 int
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	var num2 int
	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	print(sum(num1, num2))

}

func sum(x, y int) int {
	return x + y
}

//Escreva uma função que receba dois números inteiros como parâmetros e retorne a soma deles.
