package main

import "fmt"

type Car struct {
	brand string
	model string
	year  int
}

func main() {

	celta := Car{brand: "Chevrolet", model: "Celta", year: 2015}
	ka := Car{brand: "Ford", model: "Ka", year: 2018}
	up := Car{brand: "Volkswagen", model: "Up", year: 2010}
	hb20 := Car{brand: "Hyundai", model: "HB20", year: 2013}

	cars := []Car{celta, ka, up, hb20}

	carsLessThan2015 := carsWithYearLessThan2015(cars)
	for i := 0; i <= len(carsLessThan2015)-1; i++ {
		fmt.Printf("Marca=%s Modelo=%s Ano=%d \n", carsLessThan2015[i].brand, carsLessThan2015[i].model,
			carsLessThan2015[i].year)
	}

}

func carsWithYearLessThan2015(cars []Car) []Car {
	var carsLessThan2015 = []Car{}
	for i := 0; i <= len(cars)-1; i++ {
		if cars[i].year <= 2015 {
			carsLessThan2015 = append(carsLessThan2015, cars[i])
		}
	}

	return carsLessThan2015
}

//Crie uma struct `Car` com os campos `brand`, `model` e `year`. Implemente uma função que receba uma lista de `Car` e retorne apenas os carros fabricados a partir de 2015.
