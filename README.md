# System Monitoring Daemon

### Features:
* GRPC server
* In-memory storage
* Supported OS: Mac OS, Linux
* Concurrency for snapshots

### Run daemon
``go run cmd/daemon/main.go``

``go run cmd/daemon/main.go -port=50000``

#### Server Flags:
* port - GRPC server port (default 50005)

### Run client
``go run cmd/client/main.go``

``go run cmd/client/main.go -port=50000``

``go run cmd/client/main.go -port=50000 -n=1 -m=5``

#### Client lags:
* port - GRPC server port (default 50005)
* n - getting stats frequency (in sec, default 5)
* m - average stats interval (in sec, default 15)

### Build daemon

``go build -v -o ./bin/daemon ./cmd/daemon && ./bin/daemon -port=5000``

### Build client

``go build -v -o ./bin/client ./cmd/client && ./bin/client -port=5000 -n=1 -m=5``

### Configs
Are available in ``configs/config.yaml``

Set true/false for stats items.

Config example:
```
stats:
  loadavg: true
  cpu: false
  disk: false
  nettop: false
  netstat: false
```
