package gateway

import (
	"log"

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
		}
	}
	*reply = true
	return nil
}

type remoteProxy struct {
	peer string
}

func (p *remoteProxy) Error(err error) {
	log.Println("plugin["+p.peer+"]:", err)
}

/*func (p *remoteProxy) FilterRequest(req *http.Request) (bool, string, int) {

}*/

func (p *remoteProxy) Images() []demo.Image {
	var images []demo.Image
	call, err := Peer.OpenCall(p.peer, "Plugin.Images", nil, &images)
	if err != nil {
		p.Error(err)
		return []demo.Image{}
	}
	<-call.Done
	return images
}
