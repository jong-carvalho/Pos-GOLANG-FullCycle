package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Square struct {
	side float64
}

func main() {

	var choose string
	fmt.Print("Deseja ver o raio de um quadrado ou circulo? ")
	fmt.Scanln(&choose)

	if choose == "circulo" {
		var radius float64
		fmt.Print("Digite o raio do circulo: ")
		fmt.Scanln(&radius)

		circle := Circle{radius}

		fmt.Printf("A área do circulo é: %.2f", circle.Area())

	} else {
		var side float64
		fmt.Print("Digite o lado do quardado: ")
		fmt.Scanln(&side)

		square := Square{side}

		fmt.Printf("A área do quadrado é: %.2f", square.Area())
	}

}

func (circle Circle) Area() float64 {
	area := math.Pi * math.Pow(circle.radius, 2)
	return area
}

func (square Square) Area() float64 {
	return math.Pow(square.side, 2)
}

//Crie uma interface `Shape` com o método `Area()`. Implemente essa interface em duas structs: `Circle` ( com o campo `Radius`) e `Square` (com o campo side).
