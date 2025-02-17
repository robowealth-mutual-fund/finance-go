package resty

import (
	//"account-service/internals/config"
	//"account-service/internals/infrastructure/http_client"
	"github.com/robowealth-mutual-fund/finance-go/http_client"
	//httpClient "github.com/robowealth-mutual-fund/finance-go/http_client"
)

type Repository struct {
	restyClient *httpClient.Client
}
