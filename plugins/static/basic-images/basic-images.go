package images

import (
	"github.com/progrium/plugin-demo/demo"
)

func init() {
	demo.ImageProviders.Register(new(BasicImages))
}

type BasicImages struct{}

func (p *BasicImages) Images() []demo.Image {
	return []demo.Image{
		demo.Image{
			ID:   "762g3yh4n",
			Name: "ubuntu",
		},
		demo.Image{
			ID:   "86ag3y2dn",
			Name: "fedora",
		},
		demo.Image{
			ID:   "123f3yhee",
			Name: "busybox",
		},
	}
}
