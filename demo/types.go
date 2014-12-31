package demo

import (
	"net/http"
)

// interfaces

type RequestFilter interface {
	FilterRequest(req *http.Request) (pass bool, err string, status int)
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
