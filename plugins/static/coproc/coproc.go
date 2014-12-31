package coproc

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/progrium/plugin-demo/demo/gateway"
	"github.com/progrium/plugn/coproc"
)

func init() {
	log.Println("starting plugins...")
	go gateway.Serve("unix:///tmp/demo-gateway.sock")
	go coproc.StartCoprocs(findPlugins(), os.Stdout, "coproc")
}

func findPlugins() []string {
	var plugins []string
	found, err := ioutil.ReadDir("plugins/coproc")
	if err != nil {
		log.Println(err)
		return []string{}
	}
	for _, file := range found {
		filepath := "plugins/coproc/" + file.Name() + "/" + file.Name()
		if _, err := os.Stat(filepath); err == nil {
			plugins = append(plugins, filepath)
		}
	}
	return plugins
}
