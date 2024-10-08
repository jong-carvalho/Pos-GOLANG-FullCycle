package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero int
	Cidade string
	Estado string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	//compondo
	Endereco
//	criando variavel do tipo endereco
	Address Endereco
}

func (c Cliente) desativar(){
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func main() {

	jonatas := Cliente{
		Nome:  "Jonatas",
		Idade: 33,
		Ativo: true,
	}

	//é a mesma coisa
	jonatas.Cidade = "São Paulo"
	jonatas.Endereco.Cidade = "São Paulo"

	jonatas.desativar()

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", jonatas.Nome, jonatas.Idade, jonatas.Ativo)

}
