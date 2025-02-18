package resty

import (
	"github.com/robowealth-mutual-fund/finance-go/http_client"
)

type Repository struct {
	restyClient *httpClient.Client
}
