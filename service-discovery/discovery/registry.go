package discovery

import (
	"context"

	"go.etcd.io/etcd/clientv3"
)

func Register(host string) error {
	_, err := etcdClient.Put(context.Background(), etcdPrefix+":"+host, "1")
	return err
}

func Unregister(host string) error {
	_, err := etcdClient.Delete(context.Background(), etcdPrefix+":"+host)
	return err
}

func listAllHosts() ([]string, error) {
	response, err := etcdClient.Get(context.Background(), etcdPrefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	var hosts []string
	for _, kv := range response.Kvs {
		hosts = append(hosts, convertValueToHost(string(kv.Key)))
	}
	return hosts, nil
}
