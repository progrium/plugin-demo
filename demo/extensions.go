package demo

import (
	"github.com/progrium/plugin-demo/demo/extensions"
)

// RequestFilter

var RequestFilters = &requestFilters{
	extensions.NewExtensionPoint(new(RequestFilter)),
}

type requestFilters struct {
	*extensions.ExtensionPoint
}

func (ep *requestFilters) All() []RequestFilter {
	all := make([]RequestFilter, 0)
	for _, v := range ep.ExtensionPoint.All() {
		all = append(all, v.(RequestFilter))
	}
	return all
}

// RequestHandler

var RequestHandlers = &requestHandlers{
	extensions.NewExtensionPoint(new(RequestHandler)),
}

type requestHandlers struct {
	*extensions.ExtensionPoint
}

func (ep *requestHandlers) All() []RequestHandler {
	all := make([]RequestHandler, 0)
	for _, v := range ep.ExtensionPoint.All() {
		all = append(all, v.(RequestHandler))
	}
	return all
}

// ImageProvider

var ImageProviders = &imageProviders{
	extensions.NewExtensionPoint(new(ImageProvider)),
}

type imageProviders struct {
	*extensions.ExtensionPoint
}

func (ep *imageProviders) All() []ImageProvider {
	all := make([]ImageProvider, 0)
	for _, v := range ep.ExtensionPoint.All() {
		all = append(all, v.(ImageProvider))
	}
	return all
}
