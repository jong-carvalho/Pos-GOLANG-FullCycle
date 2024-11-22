package main

import "fmt"

func main() {

	countries := map[string]string{
		"Brasil":    "Brasilia",
		"Argentina": "Buenos Aires",
		"USA":       "Washington",
	}

	var country string
	fmt.Print("Digite o nome do país: ")
	fmt.Scanln(&country)

	result, err := searchCapital(country, countries)

	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println(result)
	}
}

func searchCapital(country string, countries map[string]string) (string, error) {

	for key, value := range countries {
		if country == key {
			return "A capital de " + key + " é " + value, nil
		}
	}
	return "", fmt.Errorf("%s não foi encontrado", country)
}

//Crie um map que mapeie nomes de países para suas respectivas capitais. Escreva uma função que receba o nome de um país e retorne sua capital, ou uma mensagem de erro caso o país não esteja no map.
