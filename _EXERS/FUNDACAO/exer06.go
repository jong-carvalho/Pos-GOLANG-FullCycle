package main

import "fmt"

func main() {

	fruits := []string{"morango", "manga", "kiwi"}
	//newFruits := []string{}

	for {

		var fruit string
		fmt.Print("Digite o nome da fruta ou sair para sair: ")
		fmt.Scanln(&fruit)

		if fruit == "sair" {
			break
		} else {
			fruits = addFruits(fruit, fruits)
		}
	}

	for i := 0; i < len(fruits); i++ {
		print(fruits[i] + " ")
	}

}

func addFruits(fruit string, fruits []string) []string {

	fruits = append(fruits, fruit)

	return fruits
}

//Crie um slice de strings contendo nomes de frutas. Escreva uma função que adicione mais nomes ao slice e retorne o slice atualizado.
