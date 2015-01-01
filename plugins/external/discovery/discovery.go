package main

import (
	"log"
	"os"

	"github.com/progrium/duplex/poc2/duplex"
	"github.com/progrium/duplex/poc2/duplex/rpc"
)

type RegisterArgs struct {
	Name       string
	Interfaces []string
}

type Plugin struct{}

func (p *Plugin) MatchEndpoint(_ interface{}, reply *map[string]string) error {
	*reply = map[string]string{
		"method": "GET",
		"path":   "/discovery",
	}
	return nil
}

func (p *Plugin) Handle(_ interface{}, reply *map[string][]string) error {
	*reply = map[string][]string{
		"IPs": []string{"10.0.0.1", "10.0.0.2", "10.0.0.3"},
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
	plugin.SetOption(duplex.OptName, "discovery")
	log.Println("connecting to", endpoint, "...")
	err := plugin.Connect(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	success := false
	err = plugin.Call("PluginGateway.Register",
		&RegisterArgs{"discovery", []string{"RequestHandler"}}, &success)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("registered")
	plugin.Register(new(Plugin))
	plugin.Serve()
}
