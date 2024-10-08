package main

import "fmt"

var (
	b bool
	c int
	d string
	e float64
)

type ID int

var a ID = 1

func main() {
	fmt.Printf("O tipo de B é %T", b)
	println()
	fmt.Printf("O tipo de C é %T", c)
	println()
	fmt.Printf("O tipo de D é %T", d)
	println()
	fmt.Printf("O tipo de E é %T", e)
	println()
	fmt.Printf("O tipo de A é %T", a)
	println()
	fmt.Printf("O valor de A é %v", a)
}
