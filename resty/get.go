package resty

import (
	"context"
	"fmt"
	"net/http"

	"github.com/robowealth-mutual-fund/finance-go/models/resty"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

const (
	ErrGetData = "Error get data"
	ErrIsNotOK = "Error is not 200 ok"
)

func (r *Repository) Get(host, path string, header map[string]string, pathParams, queryParams map[string]string, entity interface{}) error {
	response, err := r.restyClient.RequestWithJSON(&resty.Request{
		Header:      header,
		Host:        host,
		Method:      resty.MethodGet,
		Path:        path,
		PathParams:  pathParams,
		QueryParams: queryParams,
	})
	if err != nil {
		return errors.Wrap(err, ErrGetData)
	}

	body := response.Body()
	statusCode := response.StatusCode()

	if statusCode != http.StatusOK {
		errorMessage := fmt.Sprintf("%s %s", response.Status(), string(response.Body()))
		return errors.Wrap(errors.New(errorMessage), ErrIsNotOK)
	}

	return unmarshal(body, &entity)
}

func (r *Repository) Post(ctx context.Context, header, queryParams, formParams map[string]string, host, path string, request, entity interface{}) error {
	response, err := r.restyClient.RequestWithJSON(&resty.Request{
		Header:      header,
		Host:        host,
		Method:      resty.MethodPost,
		Path:        path,
		Body:        request,
		FormParams:  formParams,
		QueryParams: queryParams,
	})
	if err != nil {
		return errors.Wrap(err, ErrGetData)
	}

	body := response.Body()
	statusCode := response.StatusCode()

	if statusCode != http.StatusOK && statusCode != http.StatusCreated {
		errorMessage := fmt.Sprintf("%s %s", response.Status(), string(response.Body()))
		return errors.Wrap(errors.New(errorMessage), ErrIsNotOK)
	}

	return unmarshal(body, &entity)
}

func unmarshal(data []byte, v interface{}) error {
	return jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, v)
}
