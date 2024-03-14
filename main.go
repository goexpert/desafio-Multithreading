package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/goexpert/desafio-Multithreading/adapter"
)

var cep = "02206000"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client := adapter.NewClientHttp(ctx)
	urlBrasilApi := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	urlViaCep := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	body, err := client.GetRequest(urlBrasilApi)
	if err != nil {
		log.Fatal(err.Error())
	}
	println(string(body))

	body, err = client.GetRequest(urlViaCep)
	if err != nil {
		log.Fatal(err.Error())
	}
	println(string(body))

}
