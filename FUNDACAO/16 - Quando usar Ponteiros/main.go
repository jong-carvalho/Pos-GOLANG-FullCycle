package main

func soma(a, b int) int {
	a = 50
	return a + b
}

func somaAlterandoMemoria(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {

	minhaVar1 := 10
	minhaVar2 := 20

	//a função usa cópias dos valores armazenados em memória
	println(soma(minhaVar1, minhaVar2))
	println(minhaVar1)

	//usando ponteiros
	println(somaAlterandoMemoria(&minhaVar1, &minhaVar2))
	println(minhaVar1)

}
