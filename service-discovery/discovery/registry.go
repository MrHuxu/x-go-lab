package discovery

import (
	"context"
)

func Register(host string) error {
	_, err := etcdClient.Put(context.Background(), etcdPrefix+":"+host, "1")
	return err
}

func Unregister(host string) error {
	_, err := etcdClient.Delete(context.Background(), etcdPrefix+":"+host)
	return err
}
