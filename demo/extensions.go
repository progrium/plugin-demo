package demo

import (
	"log"
	"path"
	"runtime"
)

// RequestFilter

var RequestFilters = new(requestFilterExt)

type requestFilterExt struct {
	filters []RequestFilter
}

func (ext *requestFilterExt) All() []RequestFilter {
	return ext.filters
}

func (ext *requestFilterExt) Register(filter RequestFilter) {
	if ext.filters == nil {
		ext.filters = make([]RequestFilter, 0)
	}
	_, file, _, _ := runtime.Caller(1)
	log.Println("registering RequestFilter via", path.Base(file))
	ext.filters = append(ext.filters, filter)
}

// ImageProvider

var ImageProviders = new(imageProviderExt)

type imageProviderExt struct {
	providers []ImageProvider
}

func (ext *imageProviderExt) All() []ImageProvider {
	return ext.providers
}

func (ext *imageProviderExt) Register(provider ImageProvider) {
	if ext.providers == nil {
		ext.providers = make([]ImageProvider, 0)
	}
	_, file, _, _ := runtime.Caller(1)
	log.Println("registering ImageProvider via", path.Base(file))
	ext.providers = append(ext.providers, provider)
}
