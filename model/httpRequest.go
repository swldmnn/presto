package model

import "fmt"

type HttpRequest struct {
	Method  string
	Url     string
	Headers map[string]string
	Body    string
}

func NewHttpRequest(method string, url string) HttpRequest {
	return HttpRequest{method, url, make(map[string]string), ""}
}

func (r HttpRequest) String() string {
	return fmt.Sprintf("%s %s", r.Method, r.Url)
}
