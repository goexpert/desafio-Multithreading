package adapter

import (
	"context"
	"errors"
	"io"
	"net/http"
)

var (
	errTimeout = errors.New("timeout")
	errGeneral = errors.New("general")
)

type ClientHttp struct {
	Ctx context.Context
}

func NewClientHttp(ctx context.Context) *ClientHttp {
	return &ClientHttp{
		Ctx: ctx,
	}
}

func (c *ClientHttp) GetRequest(url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(c.Ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if c.Ctx.Err() == context.DeadlineExceeded {
			return nil, errTimeout
		} else {
			return nil, errGeneral
		}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
