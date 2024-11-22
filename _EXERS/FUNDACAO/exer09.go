package main

func main() {
	print(sumOfNumbers(2, 5, 8, 10))
}

func sumOfNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

//Crie uma função variádica que receba uma lista de números inteiros e retorne a soma de todos eles.
