package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	//c := http.Client{Timeout: time.Second}
	c := http.Client{Timeout: time.Microsecond}
	resp, err := c.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
