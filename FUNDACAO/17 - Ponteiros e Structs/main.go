package main

import "fmt"

type Cliente struct {
	nome string
}

func (c Cliente) andou() {
	c.nome = "Jonatas Carvalho"
	fmt.Printf("O cliente %v andou", c.nome)
}

func main() {
	jonatas := Cliente{nome: "Jonatas"}

	jonatas.andou()
	println()
	fmt.Printf("O valor da struct com o nome %v", jonatas.nome)
}
