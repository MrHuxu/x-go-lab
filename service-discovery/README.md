# Service Discovery

An experimental project about service discovery.

## Dependencies

1. [etcd](https://coreos.com/etcd/): Distributed key-value system, used as registry center.
2. [consistent](https://github.com/stathat/consistent): Implementation of [Consistent Hashing Algorithm](https://en.wikipedia.org/wiki/Consistent_hashing), to achieve load balance.

## Submodules

1. `proto`: Service definitions.
2. `discovery`: Implementation of service discovery.
3. `server`: Launch and register/unregister services, the program has two flags:
      - `port`: Port of the TCP server.
      - `sign`: Signature of the service.
4. `client`: Launch several clients to request registerd servers concurrently.

## Usage

1. Launch `etcd` and make it run on port `2379`.

2. Enter directory `client`.

        go run main.go

3. Enter directory `server` and lauch several instances.

        go run main.go -port=1234 -sign=1
        go run main.go -port=1235 -sign=2
        go run main.go -port=1236 -sign=3

After all of the steps above, just check the output of client, it will show you the RPC results, which are returned by the services **discovered** by the `discovery` module.

![](https://raw.githubusercontent.com/MrHuxu/x-go-lab/master/service-discovery/service%20discovery.gif)