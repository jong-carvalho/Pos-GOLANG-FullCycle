package main

import (
	"fmt"
	"github.com/jong-carvalho/curso-go/matematica"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Printf("Resultado: %v", s)
}
