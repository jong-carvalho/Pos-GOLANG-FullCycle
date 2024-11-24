package main

import "fmt"

func identifyType(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("O valor é um inteiro: %d\n", v)
	case string:
		fmt.Printf("O valor é uma string: %s\n", v)
	case bool:
		fmt.Printf("O valor é um booleano: %t\n", v)
	default:
		fmt.Println("Tipo não reconhecido.")
	}
}

func main() {
	identifyType("string")
	identifyType(45)
	identifyType(false)
}

//Escreva uma função que receba um parâmetro do tipo `interface{}`. Dependendo do tipo do valor recebido (inteiro, string ou booleano), imprima uma mensagem diferente.
