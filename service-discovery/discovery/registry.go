package discovery

import (
	"context"
)

// Register registers a host as a service
func Register(host string) error {
	_, err := etcdClient.Put(context.Background(), etcdPrefix+":"+host, "1")
	return err
}

// Unregister unregisters a service by its host
func Unregister(host string) error {
	_, err := etcdClient.Delete(context.Background(), etcdPrefix+":"+host)
	return err
}
