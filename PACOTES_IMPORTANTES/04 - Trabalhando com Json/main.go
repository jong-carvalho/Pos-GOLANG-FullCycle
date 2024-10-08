package main

import (
	"encoding/json"
	"os"
)

// essas anotações ajudam na hora de fazer o mapeamento dos campos
type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {

	// quando utilizamos o Marshal estamos guardando o valor do Json
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)

	if err != nil {
		println(err)
	}

	println(string(res))

	//	quando utilizamos o Encode quer dizer que estamos convertendo e ja entregando para outra pessoa
	err = json.NewEncoder(os.Stdout).Encode(conta)

	if err != nil {
		println(err)
	}

	//	o UnMarshal faz o caminho contrario ou seja estamos transformando um json em conta
	jsonPuro := []byte(`{"n": 2, "s":200}`)
	var contaX Conta

	err = json.Unmarshal(jsonPuro, &contaX)

	if err != nil {
		println(err)
	}

	println(contaX.Numero, contaX.Saldo)
}
