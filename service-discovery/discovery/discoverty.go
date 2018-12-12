package discovery

import (
	"context"
	"log"
	"sync"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
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
	d.refreshEndpoints()
	go d.monitorEndpoints()

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

func (d *discovery) refreshEndpoints() {
	response, err := etcdClient.Get(context.Background(), etcdPrefix, clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}

	for _, kv := range response.Kvs {
		d.addEndpoint(convertKeyToHost(string(kv.Key)))
	}
}

func (d *discovery) monitorEndpoints() {
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

func (d *discovery) delEndpoint(host string) error {
	d.consistent.Remove(host)

	d.lock.Lock()
	defer d.lock.Unlock()
	delete(d.mapHostToEndpoint, host)

	return nil
}
