package main

import (
	"github.com/progrium/plugin-demo/demo"
	//"github.com/progrium/plugin-demo/demo/gateway"
)

func main() {
	//go gateway.Serve("tcp://127.0.0.1:9888")
	demo.Run()
}
