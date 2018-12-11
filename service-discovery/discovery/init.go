package discovery

import (
	"log"
	"strings"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	etcdClient *clientv3.Client
	etcdPrefix = "service-discovery"
)

const ()

func init() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	etcdClient = cli
}

func convertValueToHost(value string) string {
	return strings.Split(value, etcdPrefix+":")[1]
}
