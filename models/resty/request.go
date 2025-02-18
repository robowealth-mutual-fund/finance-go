package resty

type Method string

const (
	MethodGet    Method = "GET"
	MethodPost   Method = "POST"
	MethodPut    Method = "PUT"
	MethodPatch  Method = "PATCH"
	MethodDelete Method = "DELETE"
)

type Request struct {
	Header      map[string]string
	Host        string
	Method      Method
	Path        string
	PathParams  map[string]string
	QueryParams map[string]string
	FormParams  map[string]string
	Body        interface{}
}

type Response struct {
	Status     string
	StatusCode int
	Body       interface{}
}
