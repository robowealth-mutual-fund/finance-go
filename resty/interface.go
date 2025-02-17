package resty

import "context"

//go:generate mockery --name=Interface
type Interface interface {
	Get(host, path string, header map[string]string, pathParams, queryParams map[string]string, entity interface{}) error
	Post(ctx context.Context, header, queryParams, formParams map[string]string, host, path string, request, entity interface{}) error
}
