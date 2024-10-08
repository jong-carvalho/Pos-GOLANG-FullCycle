package main

/**/
func somaInteiro(m map[string]int) int {
	var soma int

	//aqui está ignorando a string utilizando o _ e somando apenas o inteiro
	for _, v := range m {
		soma += v
	}
	return soma
}

func somaFloat(m map[string]float64) float64 {
	var soma float64

	//aqui está ignorando a string utilizando o _ e somando apenas o inteiro
	for _, v := range m {
		soma += v
	}
	return soma
}

// podemos utilizar algo desse tipo porém ele dá erro, pois embora myNumber seja um inteiro ele não é reconhecido como inteiro
type myNumber int

// por isso colocamos o ~ para que outros tipos que sejam do tipo int e float sejam reconhecidos
type Number interface {
	~int | ~float64
}

//da pra fazer assim ou utilizar as constraints
//func soma[T int | float64] (m map[string]T) T {

func soma[T Number](m map[string]T) T {
	var soma T

	//aqui está ignorando a string utilizando o _ e somando apenas o inteiro
	for _, v := range m {
		soma += v
	}
	return soma
}

//utilizando o any ele reclama pois ele não tem certeza de que esse dois operandos sao do mesmo tipo para serem comparados
//func comprara[T any](a T, b T) bool {

// com o comparable conseguimos comparar (desde que os tipos sejam iguais)
func compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Joao": 1000, "Julia": 2000, "Maria": 3000}
	m2 := map[string]float64{"Joao": 1000.50, "Julia": 2000.75, "Maria": 3000.0}
	println(somaInteiro(m))
	println(somaFloat(m2))
	//	utilizar esse contexto acima é ruim pois temos que criar duas funções que fazem a mesma coisa, só muda o tipo

	//com isso podemos utilizar as generics que normalmente são representadas pelo T
	println(soma(m))
	println(soma(m2))

	m3 := map[string]myNumber{"Joao": 1000, "Julia": 2000, "Maria": 3000}
	print(soma(m3))

	println(compara(10, 10))
}
