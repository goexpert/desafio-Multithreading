package main

import (
	"fmt"
	"log"
	"time"

	"github.com/goexpert/desafio-Multithreading/adapter"
)

var cep = "01001000"

func getCep(client *adapter.ClientHttp, url string, resp chan []byte) {
	body, err := client.GetRequest(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	resp <- body
	close(resp)
}

func main() {
	client := adapter.NewClientHttp()
	urlBrasilApi := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	urlViaCep := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	responseBrasil := make(chan []byte)
	responseVia := make(chan []byte)

	go getCep(client, urlViaCep, responseVia)
	go getCep(client, urlBrasilApi, responseBrasil)

	select {
	case resp1 := <-responseBrasil:
		log.Printf("Received from BrasilCep %s\n", resp1)
	case resp2 := <-responseVia:
		log.Printf("Received fom ViaCep %s\n", resp2)
	case <-time.After(time.Millisecond * 1000):
		log.Println("timeout")
	}

}
