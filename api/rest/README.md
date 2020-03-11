# Rest API Service

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: hanako.api.rest
- Type: api
- Alias: rest


## Dependencies

1. This service depend on anime service, please make sure it is running.
2. This service depend on api gateway for translate http request to rpc call.
the default is using micro api gateway

    this command simplify to run the api 
    ```
    make api
    ```
3. Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.
In the event you need a resilient multi-host setup we recommend etcd.

    ```
    # install etcd
    brew install etcd
    
    # run etcd
    etcd
    ```
   
## Usage

A Makefile is included for convenience

Run without compile (For development)
```cgo
make dev
```

Build the binary
```
make build
```

Run the service
```
./rest-api
```

Build a docker image
```
make docker
```