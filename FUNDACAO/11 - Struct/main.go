package main

import "fmt"

type Client struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	jonatas := Client{
		Nome:  "Jonatas",
		Idade: 33,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", jonatas.Nome, jonatas.Idade, jonatas.Ativo)

}
