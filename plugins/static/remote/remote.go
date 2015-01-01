package remote

import (
	"log"

	"github.com/progrium/plugin-demo/demo/gateway"
)

func init() {
	log.Println("listening for plugins...")
	go gateway.Serve("tcp://127.0.0.1:9888")
}
