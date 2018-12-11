package discovery

import (
	"log"
	"sync"

	"stathat.com/c/consistent"
)

type Discovery interface {
	Get(string) Endpoint
}

func New(fn newEndpointFunc) Discovery {
	d := &discovery{
		newEndpointFunc:   fn,
		mapHostToEndpoint: make(map[string]Endpoint),
		consistent:        consistent.New(),
	}

	hosts, err := listAllHosts()
	if err != nil {
		log.Fatal(err)
	}
	for _, host := range hosts {
		d.addEndpoint(host)
	}

	return d
}

type newEndpointFunc func(host string) Endpoint

type discovery struct {
	newEndpointFunc   newEndpointFunc
	mapHostToEndpoint map[string]Endpoint
	consistent        *consistent.Consistent
	lock              sync.RWMutex
}

func (d *discovery) Get(key string) Endpoint {
	d.consistent.Get(key)
	host, err := d.consistent.Get(key)
	if err != nil {
		log.Fatal(err)
	}

	if _, ok := d.mapHostToEndpoint[host]; !ok {
		log.Fatal("endpoint not found")
	}
	return d.mapHostToEndpoint[host]
}

func (d *discovery) addEndpoint(host string) {
	d.consistent.Add(host)

	d.lock.Lock()
	defer d.lock.Unlock()
	d.mapHostToEndpoint[host] = d.newEndpointFunc(host)
}

func (d *discovery) delEndpoint(host string) error {
	d.consistent.Remove(host)

	d.lock.Lock()
	defer d.lock.Unlock()
	delete(d.mapHostToEndpoint, host)

	return nil
}
