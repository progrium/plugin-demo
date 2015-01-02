package coproc

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/progrium/plugin-demo/demo/coproc"
	"github.com/progrium/plugin-demo/demo/gateway"
)

func init() {
	log.Println("starting plugins...")
	log.Println(os.Getpid())
	go gateway.Serve("unix:///tmp/demo-gateway.sock")
	host := coproc.StartHost(findPlugins())
	go func() {
		handler := make(chan os.Signal, 1)
		signal.Notify(handler, os.Interrupt)
		first := true
		for sig := range handler {
			switch sig {
			case os.Interrupt:
				log.Println("ctrl-c detected")
				host.Shutdown(!first)
				go func() {
					host.Wait()
					os.Exit(0)
				}()
				first = false
			}
		}
	}()
}

func findPlugins() []string {
	var plugins []string
	found, err := ioutil.ReadDir("plugins/external")
	if err != nil {
		log.Println(err)
		return []string{}
	}
	for _, file := range found {
		filepath := "plugins/external/" + file.Name() + "/" + file.Name()
		if _, err := os.Stat(filepath); err == nil {
			plugins = append(plugins, filepath)
		}
	}
	return plugins
}
