package main

// go get github.com/google/uuid

func main() {

	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "três", "quatro"}

	for k, v := range numeros {
		print(k)
		print("-")
		print(v)
		print(" ... ")
	}

	println()

	for _, numero := range numeros {
		println(numero)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	//loop infinito
	for {
		println("Hello World")
	}
}
