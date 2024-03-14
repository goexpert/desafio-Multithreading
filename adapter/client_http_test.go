package adapter

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBrasilAPI(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := NewClientHttp(ctx)
	body, err := client.GetRequest("https://brasilapi.com.br/api/cep/v1/01153000")
	assert.Nil(t, err)
	assert.NotNil(t, body)
}

func TestViaCep(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := NewClientHttp(ctx)
	body, err := client.GetRequest("http://viacep.com.br/ws/01153000/json/")
	assert.Nil(t, err)
	assert.NotNil(t, body)
}
