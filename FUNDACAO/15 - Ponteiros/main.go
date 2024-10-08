package main

func main() {

	//memória -> endereço -> valor
	a := 10
	println(a)
	//endereço de memória
	println(&a)

	//quando coloca o asterisco se refere ao endereço
	var ponteiro *int = &a
	println(ponteiro)

	//altera o valor que tem armazenado na memória
	*ponteiro = 20
	println(a)

	b := &a
	println(b)
	println(*b)

}
