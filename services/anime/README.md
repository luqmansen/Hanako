# Anime Service

This is the Anime service

Generated with

```
micro new anime --namespace=hanako --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Requirements](#requirements)
- [Dependencies](#dependencies)
- [Setup](#setup)
- [Usage](#usage)

## Configuration

- FQDN: hanako.srv.anime
- Type: srv
- Alias: anime

## Requirements
- Mongodb:
<br>specify mongodb url and database name on .env.


For development purpose, I personally use docker.
I made custom yml to run some of container that needed on this services. 

Run this command
```cgo
make du # For run mongodb
make dd # For take down
```

## Setup
1. Create a ``.env`` file on the folder
2. Setup the dev environment there, format
```dotenv
mongo_url=mongodb_url
mongo_dbname=mongo_db_database_name
```

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./anime-srv
```

Build a docker image
```
make docker
```