package main

import (
	"log"

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
	plugin := rpc.NewPeer()
	defer plugin.Shutdown()
	plugin.SetOption(duplex.OptName, "scripting")
	log.Println("connecting...")
	err := plugin.Connect("unix:///tmp/demo-gateway.sock")
	if err != nil {
		log.Fatal(err)
	}
	args := &RegisterArgs{"scripting", []string{"ImageProvider"}}
	success := false
	err = plugin.Call("PluginGateway.Register", args, &success)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("registered")
	plugin.Register(new(Plugin))
	plugin.Serve()
}
