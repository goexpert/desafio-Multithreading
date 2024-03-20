package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/goexpert/desafio-Multithreading/adapter"
)

var cep = "02206000"

func getCep(client *adapter.ClientHttp, url string, wg *sync.WaitGroup) ([]byte, error) {
	body, err := client.GetRequest(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	println(body)
	wg.Done()
	return body, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client := adapter.NewClientHttp(ctx)
	urlBrasilApi := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	urlViaCep := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	go getCep(client, urlBrasilApi, &waitGroup)
	go getCep(client, urlViaCep, &waitGroup)

	waitGroup.Wait()

}
