package gateway

import (
	"log"
	"net/http"
	"net/url"

	"github.com/progrium/duplex/poc2/duplex/rpc"
	"github.com/progrium/plugin-demo/demo"
)

var Peer *rpc.Peer

func Serve(endpoint string) {
	Peer = rpc.NewPeer()
	defer Peer.Shutdown()
	err := Peer.Bind(endpoint)
	if err != nil {
		panic(err)
	}
	Peer.Register(new(PluginGateway))
	Peer.Serve()
}

type PluginGateway struct{}

type RegisterArgs struct {
	Name       string
	Interfaces []string
}

func (g *PluginGateway) Register(args *RegisterArgs, reply *bool) error {
	for _, iface := range args.Interfaces {
		switch iface {
		case "ImageProvider":
			demo.ImageProviders.Register(&remoteProxy{args.Name})
		case "RequestHandler":
			demo.RequestHandlers.Register(&remoteProxy{args.Name})
		}
	}
	*reply = true
	return nil
}

type remoteProxy struct {
	peer string
}

func (p *remoteProxy) call(method string, input, output interface{}) error {
	call, err := Peer.OpenCall(p.peer, method, input, output)
	if err != nil {
		log.Println("plugin["+p.peer+"]:", err)
		return err
	}
	<-call.Done
	return nil
}

/*func (p *remoteProxy) FilterRequest(req *http.Request) (bool, string, int) {

}*/

func (p *remoteProxy) MatchEndpoint() (method string, path string) {
	var endpoint map[string]string
	err := p.call("Plugin.MatchEndpoint", nil, &endpoint)
	if err != nil {
		return "", ""
	}
	return endpoint["method"], endpoint["path"]
}

func (p *remoteProxy) Handle(u *url.URL, h http.Header, input interface{}) (int, http.Header, interface{}, error) {
	var obj map[string][]string
	err := p.call("Plugin.Handle", nil, &obj)
	if err != nil {
		return http.StatusInternalServerError, nil, nil, nil
	}
	return http.StatusOK, nil, obj, nil
}

func (p *remoteProxy) Images() []demo.Image {
	var images []demo.Image
	err := p.call("Plugin.Images", nil, &images)
	if err != nil {
		return []demo.Image{}
	}
	return images
}
