package demo

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/rcrowley/go-tigertonic"
)

func marshal(obj interface{}) []byte {
	bytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Println("marshal:", err)
	}
	return bytes
}

func unmarshal(input []byte, obj interface{}) error {
	err := json.Unmarshal(input, obj)
	if err != nil {
		return err
	}
	return nil
}

func Run() {
	mux := tigertonic.NewTrieServeMux()
	mux.Handle("GET", "/hello",
		tigertonic.Marshaled(get_hello),
	)
	mux.Handle("GET", "/images",
		tigertonic.Marshaled(get_images),
	)
	log.Println("listening on :8000...")
	tigertonic.NewServer(":8000",
		tigertonic.ApacheLogged(
			tigertonic.If(applyRequestFilters, mux))).ListenAndServe()
}

func applyRequestFilters(req *http.Request) (http.Header, error) {
	for _, filter := range RequestFilters.All() {
		ok, err, status := filter.FilterRequest(req)
		if !ok {
			return http.Header{}, tigertonic.NewHTTPEquivError(errors.New(err), status)
		}
	}
	return nil, nil
}

type HelloResponse struct {
	Text  string
	Extra int
}

func get_hello(u *url.URL, h http.Header, _ interface{}) (int, http.Header, *HelloResponse, error) {
	return http.StatusOK, nil, &HelloResponse{"Hello world", 100}, nil
}

func get_images(u *url.URL, h http.Header, _ interface{}) (int, http.Header, []Image, error) {
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
