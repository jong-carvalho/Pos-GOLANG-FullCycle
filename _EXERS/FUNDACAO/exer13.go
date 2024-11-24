package main

import "fmt"

type Rectangle struct {
	width  float32
	height float32
}

func main() {
	var width float32
	var height float32

	fmt.Print("Digite a largura do retangulo: ")
	fmt.Scanln(&width)

	fmt.Print("Digite a altura do retangulo: ")
	fmt.Scanln(&height)

	rectangle := Rectangle{width, height}

	fmt.Printf("A área do retangulo é: %.2f", perimeter(rectangle))
}

func perimeter(rectangle Rectangle) float32 {
	return 2 * (rectangle.width + rectangle.height)
}

//Adicione um método à struct `Rectangle` que calcule e retorne o perímetro do retângulo. A struct deve ter os campos `Width` e `Height`.
