package main

import (
	"bufio"
	"fmt"
	"os"
)

type Address struct {
	street string
	city   string
}

type Employee struct {
	Address
	position string
}

func main() {

	var street string
	var city string
	var position string
	fmt.Print("Digite o nome da rua: ")
	in := bufio.NewReader(os.Stdin)
	street, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Digite o nome da sua cidade: ")
	fmt.Scanln(&city)
	fmt.Print("Digite o nome do cargo: ")
	fmt.Scanln(&position)

	address := Address{
		street,
		city,
	}
	employee := Employee{
		address,
		position,
	}
	printEmployeeInformations(employee)
}

func printEmployeeInformations(employee Employee) {
	fmt.Printf("Rua: %s, Cidade: %s, Cargo: %s", employee.street, employee.city, employee.position)
}

//Crie uma struct `Address` com os campos `Street` e `City`. Crie uma outra struct `Employee` que inclua `Address` e tenha um campo adicional chamado `Position`. Escreva uma função que receba um `Employee` e imprima todas as suas informações.
