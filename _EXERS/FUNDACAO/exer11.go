package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	var name string
	var age int
	fmt.Print("Digite o seu nome: ")
	fmt.Scanln(&name)
	fmt.Print("Digite a sua idade: ")
	fmt.Scanln(&age)

	person := Person{
		name, age,
	}

	printPerson(person)

}

func printPerson(person Person) {
	fmt.Printf("Nome: %s, Idade: %d", person.name, person.age)
}

//Crie uma struct chamada `Person` com os campos `Name` e `Age`. Escreva uma função que receba um `Person` como parâmetro e imprima suas informações.
