package main

func main() {

	a := 1
	b := 2

	if a == b {
		println("a == b")
	} else {
		println("a != b")
	}

	switch a {
	case 1:
		println("a == 1")
	case 2:
		println("a == 2")
	default:
		println("fora dos parametros")
	}
}
