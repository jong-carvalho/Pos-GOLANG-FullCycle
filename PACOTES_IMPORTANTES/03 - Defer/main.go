package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://www.google.com")

	if err != nil {
		panic(err)
	}

	//faz com que seja executado por ultimo
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

	//	EXEMPLO
	println("Primeira Linha")
	println("Segunda Linha")
	println("Terceira Linha")

	println()

	defer println("Primeira Linha")
	println("Segunda Linha")
	println("Terceira Linha")

}
