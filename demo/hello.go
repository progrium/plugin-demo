package demo

import (
	"net/http"
	"net/url"
)

type HelloHandler struct{}

func (h *HelloHandler) MatchEndpoint() (string, string) {
	return "GET", "/hello"
}

func (h *HelloHandler) Handle(u *url.URL, _ http.Header, _ interface{}) (int, http.Header, interface{}, error) {
	return http.StatusOK, nil, &map[string]string{"Text": "Hello world"}, nil
}
