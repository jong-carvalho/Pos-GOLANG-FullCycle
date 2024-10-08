package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero int
	Cidade string
	Estado string
}

type Client struct {
	Nome  string
	Idade int
	Ativo bool
	//compondo
	Endereco
//	criando variavel do tipo endereco
	Address Endereco
}

func main() {

	jonatas := Client{
		Nome:  "Jonatas",
		Idade: 33,
		Ativo: true,
	}

	//é a mesma coisa
	jonatas.Cidade = "São Paulo"
	jonatas.Endereco.Cidade = "São Paulo"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", jonatas.Nome, jonatas.Idade, jonatas.Ativo)

}
