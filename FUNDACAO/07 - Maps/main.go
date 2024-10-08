package main

import "fmt"

func main() {
	salarios := map[string]int{"Jonatas": 1000, "Daniela": 2000, "Maria": 3500}

	fmt.Println(salarios["Jonatas"])
	delete(salarios, "Jonatas")
	salarios["Wes"] = 500
	fmt.Println(salarios["Wes"])

	for nome, salario := range salarios {
		fmt.Println("O salario de %s é %d", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Println(salario)
	}
	//make cria o map vazio
	//sal := make(map[string]int)
	////também cria um map vazio
	//sal1 := map[string]int{}

}
