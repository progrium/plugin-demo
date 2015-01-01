package demo

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/rcrowley/go-tigertonic"
)

func Run() {
	RequestHandlers.Register(new(HelloHandler))
	RequestHandlers.Register(new(ImagesHandler))
	time.Sleep(1 * time.Second)

	log.Println("listening on :8000...")
	tigertonic.NewServer(":8000",
		tigertonic.ApacheLogged(
			tigertonic.If(applyRequestFilters, loadRequestHandlers()))).ListenAndServe()
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

func loadRequestHandlers() *tigertonic.TrieServeMux {
	mux := tigertonic.NewTrieServeMux()
	for _, handler := range RequestHandlers.All() {
		method, endpoint := handler.MatchEndpoint()
		mux.Handle(method, endpoint, tigertonic.Marshaled(handler.Handle))
	}
	return mux
}
