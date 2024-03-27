package adapter

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)

var (
	errTimeout = errors.New("timeout")
	errGeneral = errors.New("general")
	errNot200  = errors.New("not200")
)

type ClientHttp struct {
	Ctx context.Context
}

func NewClientHttp() *ClientHttp {
	return &ClientHttp{
		Ctx: context.Background(),
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

	if resp.StatusCode != 200 {
		time.Sleep(time.Second * 5)
		return nil, errNot200
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
