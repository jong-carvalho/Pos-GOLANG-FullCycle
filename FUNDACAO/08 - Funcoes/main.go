package main

import (
	"errors"
	"fmt"
)

func main() {

	valor, err := sum2(50, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)

	fmt.Println(sum(1, 2))

}

func sum(a int, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}
	return a + b, false
}

func sum2(a int, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
