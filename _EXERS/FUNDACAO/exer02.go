package main

import (
	"fmt"
	"strconv"
)

func main() {

	cities := map[string]int{
		"Guarulhos": 100000,
		"São Paulo": 700000,
		"Santos":    55000,
	}

	var city string
	fmt.Print("Digite o nome da cidade: ")
	fmt.Scanln(&city)

	result, err := populationOfACity(city, cities)
	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println(result)
	}

}

func populationOfACity(cityName string, cities map[string]int) (string, error) {

	for key, value := range cities {
		if cityName == key {
			return "A cidade " + key + " possui " + strconv.Itoa(value) + " habitantes", nil
		}
	}
	return "", fmt.Errorf("A cidade %s não foi encontrada", cityName)
}

//Implemente um map que armazene nomes de cidades e suas populações. Crie uma função que receba o nome de uma cidade e retorne a população.
//Se a cidade não existir, retorne um erro customizado.
