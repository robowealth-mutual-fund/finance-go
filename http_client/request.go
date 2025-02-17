package httpClient

import (
	"errors"
	"time"

	"github.com/go-resty/resty/v2"
	restyModel "github.com/robowealth-mutual-fund/finance-go/models/resty"
)

func (c *Client) RequestWithJSON(
	request *restyModel.Request,
) (*resty.Response, error) {
	var (
		response *resty.Response
		err      error
	)
	client := c.restyClient.
		SetBaseURL(request.Host).
		SetTimeout(time.Duration(300) * time.Second).
		R().
		//SetContext(*ctx).
		ForceContentType("application/json").
		SetHeaders(request.Header).
		SetBody(request.Body)

	if request.PathParams != nil {
		client.SetPathParams(request.PathParams)
	}

	if request.QueryParams != nil {
		client.SetQueryParams(request.QueryParams)
	}

	if request.FormParams != nil {
		client.SetFormData(request.FormParams)
	}

	switch request.Method {
	case resty.MethodGet:
		response, err = client.Get(request.Path)
	case resty.MethodPost:
		response, err = client.Post(request.Path)
	case resty.MethodPut:
		response, err = client.Put(request.Path)
	case resty.MethodPatch:
		response, err = client.Patch(request.Path)
	case resty.MethodDelete:
		response, err = client.Delete(request.Path)
	default:
		return nil, errors.New("method not match")
	}

	if err != nil {
		return nil, err
	}

	return response, nil
}
