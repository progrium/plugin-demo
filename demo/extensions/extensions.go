package extensions

import (
	"errors"
	"reflect"
	"sync"

	"code.google.com/p/go-uuid/uuid"
)

var ErrStrNotRegistered = "extensions: not registered: "
var ErrStrNotInterface = "extensions: cannot register extension without interface: "

var extensions = struct {
	sync.Mutex
	m map[string]*ExtensionPoint
}{
	m: make(map[string]*ExtensionPoint),
}

type ExtensionPoint struct {
	sync.Mutex
	iface      interface{}
	ifaceName  string
	active     string
	registered map[string]interface{}
}

func (ep *ExtensionPoint) Active() interface{} {
	return ep.Get(ep.active)
}

func (ep *ExtensionPoint) SetActive(name string) error {
	ep.Lock()
	defer ep.Unlock()
	_, ok := ep.registered[name]
	if !ok {
		return errors.New(ErrStrNotRegistered + name)
	}
	ep.active = name
	return nil
}

func (ep *ExtensionPoint) Get(name string) interface{} {
	ep.Lock()
	defer ep.Unlock()
	extension, ok := ep.registered[name]
	if !ok {
		return errors.New(ErrStrNotRegistered + name)
	}
	return extension
}

func (ep *ExtensionPoint) All() []interface{} {
	ep.Lock()
	defer ep.Unlock()
	all := make([]interface{}, 0)
	for _, e := range ep.registered {
		all = append(all, e)
	}
	return all
}

func (ep *ExtensionPoint) Register(extension interface{}) error {
	return ep.RegisterWithName(extension, uuid.NewRandom().String())
}

func (ep *ExtensionPoint) RegisterWithName(extension interface{}, name string) error {
	ep.Lock()
	defer ep.Unlock()
	ifaceType := reflect.TypeOf(ep.iface).Elem()
	if !reflect.TypeOf(extension).Implements(ifaceType) {
		return errors.New(ErrStrNotInterface + ifaceType.Name())
	}
	ep.registered[name] = extension
	return nil
}

func RegisterWithName(iface string, extension interface{}, name string) error {
	extensions.Lock()
	defer extensions.Unlock()
	return extensions.m[iface].RegisterWithName(extension, name)
}

func Register(iface string, extension interface{}) error {
	extensions.Lock()
	defer extensions.Unlock()
	return extensions.m[iface].Register(extension)
}

func NewExtensionPoint(i interface{}) *ExtensionPoint {
	ep := &ExtensionPoint{
		iface:      i,
		ifaceName:  reflect.TypeOf(i).Elem().Name(),
		registered: make(map[string]interface{}),
	}
	extensions.Lock()
	defer extensions.Unlock()
	extensions.m[ep.ifaceName] = ep
	return ep
}
