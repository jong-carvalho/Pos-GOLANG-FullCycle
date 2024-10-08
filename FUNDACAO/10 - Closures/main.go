package main

import "fmt"

func main() {

	total := func() int {
		return sum(1, 3, 45, 6, 89, 90, 772) * 2
	}()

	fmt.Println(sum(1, 3, 45, 6, 89, 90, 772))

	fmt.Println(total)

}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
