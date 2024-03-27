package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrasilAPI(t *testing.T) {

	client := NewClientHttp()
	body, err := client.GetRequest("https://brasilapi.com.br/api/cep/v1/01001000")
	assert.Nil(t, err)
	assert.NotNil(t, body)
}

func TestViaCep(t *testing.T) {
	client := NewClientHttp()
	body, err := client.GetRequest("http://viacep.com.br/ws/01001000/json/")
	assert.Nil(t, err)
	assert.NotNil(t, body)
}
