package basicauth

import (
	"net/http"

	"github.com/progrium/plugin-demo/demo"
)

func init() {
	demo.RequestFilters.Register(new(BasicAuthFilter))
}

var allowed = map[string]string{
	"demo": "demo",
	"jeff": "pass",
}

type BasicAuthFilter struct{}

func (f *BasicAuthFilter) FilterRequest(req *http.Request) (bool, string, int) {
	user, pass, ok := req.BasicAuth()
	if !ok {
		return false, "basic auth needed", 403
	}
	userpass, ok := allowed[user]
	if !ok {
		return false, "user not found", 403
	}
	if pass != userpass {
		return false, "password incorrect", 403
	}
	return true, "", 0
}
