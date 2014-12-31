package main

import (
	"log"
	"os"

	"github.com/progrium/duplex/poc2/duplex"
	"github.com/progrium/duplex/poc2/duplex/rpc"
	"github.com/progrium/plugin-demo/demo"
)

type RegisterArgs struct {
	Name       string
	Interfaces []string
}

type Plugin struct{}

func (p *Plugin) Images(_ interface{}, reply *[]demo.Image) error {
	*reply = []demo.Image{
		demo.Image{
			ID:   "123abc987",
			Name: "extra/doom",
		},
		demo.Image{
			ID:   "432abc111",
			Name: "extra/quake",
		},
		demo.Image{
			ID:   "654abc222",
			Name: "extra/wolf3d",
		},
	}
	return nil
}

func main() {
	endpoint := "unix:///tmp/demo-gateway.sock"
	if len(os.Args) > 1 {
		endpoint = os.Args[1]
	}
	plugin := rpc.NewPeer()
	defer plugin.Shutdown()
	plugin.SetOption(duplex.OptName, "extra-images")
	log.Println("connecting to", endpoint, "...")
	err := plugin.Connect(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	success := false
	err = plugin.Call("PluginGateway.Register",
		&RegisterArgs{"extra-images", []string{"ImageProvider"}}, &success)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("registered")
	plugin.Register(new(Plugin))
	plugin.Serve()
}
