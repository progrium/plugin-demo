package scripting

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"runtime"

	"github.com/progrium/go-scripting"
	"github.com/progrium/go-scripting/ottojs"
	"github.com/progrium/plugin-demo/demo"
)

func init() {
	_, file, _, ok := runtime.Caller(2)
	if ok && path.Base(file) == "plugins.go" {
		for name, ifaces := range LoadScripts("plugins/scripts") {
			for _, iface := range ifaces {
				switch iface {
				case "ImageProvider":
					demo.ImageProviders.Register(&scriptProxy{name})
				case "RequestFilter":
					demo.RequestFilters.Register(&scriptProxy{name})
				}
			}
		}
	}
}

func LoadScripts(path string) map[string][]string {
	ottojs.Register()
	scripting.UpdateGlobals(map[string]interface{}{
		"println": log.Println,
	})
	scripting.LoadModulesFromPath(path)
	scripts := make(map[string][]string)
	for name, _ := range scripting.Modules() {
		ret, err := scripting.Call(name, "implements", nil)
		if err != nil {
			log.Println("scripting:", err)
			continue
		}
		log.Println("Loading script", name, ret)
		scripts[name] = make([]string, 0)
		for _, iface := range ret.([]interface{}) {
			scripts[name] = append(scripts[name], iface.(string))
		}
	}
	return scripts
}

type scriptProxy struct {
	module string
}

func (p *scriptProxy) Error(err error) {
	log.Println("scripting["+p.module+"]:", err)
}

func (p *scriptProxy) FilterRequest(req *http.Request) (bool, string, int) {
	scripting.UpdateGlobals(map[string]interface{}{"req": req})
	ret, err := scripting.Call(p.module, "FilterRequest")
	if err != nil {
		p.Error(err)
		return true, "", 0
	}
	val := ret.([]interface{})
	return val[0].(bool), val[1].(string), int(val[2].(int64))
}

func (p *scriptProxy) Images() []demo.Image {
	ret, err := scripting.Call(p.module, "Images")
	if err != nil {
		p.Error(err)
		return []demo.Image{}
	}
	images := make([]demo.Image, 0)
	for _, imageMap := range ret.([]interface{}) {
		var image demo.Image
		err := mapToStruct(imageMap.(map[string]interface{}), &image)
		if err != nil {
			p.Error(err)
			continue
		}
		images = append(images, image)
	}
	return images
}

func mapToStruct(m map[string]interface{}, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}
