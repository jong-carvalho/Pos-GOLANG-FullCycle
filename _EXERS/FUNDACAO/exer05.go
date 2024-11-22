package main

import "fmt"

func main() {

	var name string
	var age int
	fmt.Print("Digite o seu nome: ")
	fmt.Scanln(&name)
	fmt.Print("Digite a sua idade: ")
	fmt.Scanln(&age)

	fmt.Printf("Meu nome é %s e tenho %d anos.", name, age)
}

//Implemente um programa que receba duas variáveis (nome e idade) e use a função `fmt.Printf` para imprimir: "Meu nome é [nome] e tenho [idade] anos.
