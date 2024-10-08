package matematica

//para utilizar funções temos que declara-las em maiusculas
func Soma[T int | float64](a, b T) T {
	return a + b
}

//a mesma coisa serve para structs
type Carro struct {
	//a mesma coisa para variaveis
	Marca string
}