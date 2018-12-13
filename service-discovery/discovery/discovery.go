package discovery

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"stathat.com/c/consistent"
)

// Discovery defines the interface that discovery needs to implement
type Discovery interface {
	Get(string) (Endpoint, error)
}

// New creates a instance of Discovery
func New(fn newEndpointFunc) Discovery {
	d := &discovery{
		newEndpointFunc:   fn,
		mapHostToEndpoint: make(map[string]Endpoint),
		consistent:        consistent.New(),
	}
	d.refreshEndpoints()
	go d.monitorRegistry()

	return d
}

type newEndpointFunc func(host string) Endpoint

type discovery struct {
	newEndpointFunc   newEndpointFunc
	mapHostToEndpoint map[string]Endpoint
	consistent        *consistent.Consistent
	lock              sync.RWMutex
}

func (d *discovery) Get(key string) (Endpoint, error) {
	host, err := d.consistent.Get(key)
	if err != nil {
		return nil, err
	}

	if _, ok := d.mapHostToEndpoint[host]; !ok {
		return nil, errors.New("endpoint not found")
	}
	return d.mapHostToEndpoint[host], nil
}

func (d *discovery) refreshEndpoints() {
	response, err := etcdClient.Get(context.Background(), etcdPrefix, clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}

	for _, kv := range response.Kvs {
		d.addEndpoint(convertKeyToHost(string(kv.Key)))
	}
}

func (d *discovery) monitorRegistry() {
	watchCh := etcdClient.Watch(context.Background(), etcdPrefix, clientv3.WithPrefix())

	for response := range watchCh {
		event := response.Events[0]

		switch event.Type {
		case mvccpb.PUT:
			d.addEndpoint(convertKeyToHost(string(event.Kv.Key)))

		case mvccpb.DELETE:
			d.delEndpoint(convertKeyToHost(string(event.Kv.Key)))
		}
	}
}

func (d *discovery) addEndpoint(host string) {
	d.consistent.Add(host)

	d.lock.Lock()
	defer d.lock.Unlock()
	d.mapHostToEndpoint[host] = d.newEndpointFunc(host)
}

func (d *discovery) delEndpoint(host string) {
	d.consistent.Remove(host)

	d.lock.Lock()
	defer d.lock.Unlock()
	delete(d.mapHostToEndpoint, host)
}
