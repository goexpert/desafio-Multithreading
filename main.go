package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/goexpert/desafio-Multithreading/adapter"
)

var cep = "02206000"

func getCep(client *adapter.ClientHttp, url string, resp chan []byte) []byte {
	body, err := client.GetRequest(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	resp <- body
	close(resp)
	// log.Println(body)
	return body
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	client := adapter.NewClientHttp(ctx)
	urlBrasilApi := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	urlViaCep := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	responseBrasil := make(chan []byte)
	responseVia := make(chan []byte)

	go getCep(client, urlBrasilApi, responseBrasil)
	go getCep(client, urlViaCep, responseVia)

	select {
	case resp1 := <-responseBrasil:
		log.Printf("Received from BrasilCep %s\n", resp1)
	case resp2 := <-responseVia:
		log.Printf("Received fom ViaCep %s\n", resp2)
	case <-time.After(time.Second * 1):
		log.Println("timeout")
	}

}
