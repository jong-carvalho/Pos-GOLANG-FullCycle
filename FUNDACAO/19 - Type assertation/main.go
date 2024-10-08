package main

import "fmt"

/**/

func main() {
	var minhaVar interface{} = "Jonatas Carvalho"

	println(minhaVar)

	println(minhaVar.(string))

	//exibe o valor em res ... e o ok é dizendo se a conversão foi bem sucedida
	//se a conversão não for bem sucedida o ok vem como false
	res, ok := minhaVar.(int)

	fmt.Println("O valor de res é %v e o resultado de ok é %v", res, ok)

	// ele vai dar erro por causa da forte tipagem do GO, o ok é obrigatório para dizer se de fato ele consegue converter a variavel
	//res2 := minhaVar.(int)
	//fmt.Printf("O valor de res2 é %v", res2)
}
