package demo

import (
	"net/http"
	"net/url"
)

// interfaces

type RequestFilter interface {
	FilterRequest(req *http.Request) (pass bool, err string, status int)
}

type RequestHandler interface {
	MatchEndpoint() (method string, path string)
	Handle(u *url.URL, h http.Header, input interface{}) (int, http.Header, interface{}, error)
}

type ImageProvider interface {
	Images() []Image
}

// data types

type Image struct {
	ID          string
	Name        string
	Description string
}
