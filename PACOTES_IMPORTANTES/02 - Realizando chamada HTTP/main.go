package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://viacep.com.br/ws/07244270/json/")

	if err != nil {
		panic(err)
	}

	res, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))
	req.Body.Close()
}
