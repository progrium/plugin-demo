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
	all_untyped := ep.ExtensionPoint.All()
	all_typed := make([]RequestHandler, len(all_untyped))
	for i, v := range all_untyped {
		all_typed[i] = v.(RequestHandler)
	}
	return all_typed
}

// ImageProvider

var ImageProviders = &imageProviders{
	extensions.NewExtensionPoint(new(ImageProvider)),
}

type imageProviders struct {
	*extensions.ExtensionPoint
}

func (ep *imageProviders) All() []ImageProvider {
	all_untyped := ep.ExtensionPoint.All()
	all_typed := make([]ImageProvider, len(all_untyped))
	for i, v := range all_untyped {
		all_typed[i] = v.(ImageProvider)
	}
	return all_typed
}
