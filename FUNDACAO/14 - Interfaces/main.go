package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero int
	Cidade string
	Estado string
}

type Pessoa interface {
	desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) desativar() {}

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

func Desativacao(pessoa Pessoa){
	pessoa.desativar()
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

	Desativacao(jonatas)

	minhaEmpresa := Empresa{}
	//qualquer struct que eu criar que tiver o método desativar vai implementar a interface
	Desativacao(minhaEmpresa)
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", jonatas.Nome, jonatas.Idade, jonatas.Ativo)

}
