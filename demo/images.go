package demo

import (
	"net/http"
	"net/url"
)

type ImagesHandler struct{}

func (h *ImagesHandler) MatchEndpoint() (string, string) {
	return "GET", "/images"
}

func (h *ImagesHandler) Handle(u *url.URL, _ http.Header, _ interface{}) (int, http.Header, interface{}, error) {
	images := []Image{
		Image{
			ID:   "123s1qt7h",
			Name: "scratch",
		},
	}
	for _, provider := range ImageProviders.All() {
		images = append(images, provider.Images()...)
	}
	return http.StatusOK, nil, images, nil
}
