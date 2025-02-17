package httpClient

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	restyClient *resty.Client
}

func New() *Client {
	return &Client{
		restyClient: resty.New(),
	}
}
