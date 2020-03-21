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
```

#### import data to mongodb for the first time
After you run the mongodb container
1. Enter to the database from the mongo client cli
    ```cgo
    $ mongo -u root -p root --authenticationDatabase admin    
    ```
2. Create new database and collection
    ```cgo
    > use hanako
    > db.createCollection("anime")
    ```
3. Enter to your database
    ```cgo
    > use hanako
    ```
4. Create new credential for your database
    ```cgo    
   > db.createUser(
       {
         user : "root",
         pwd : "root",
         roles: [ { role: 'root', db: 'admin' } ]
       }
    )
    ```     
5. Now exit, and to this on the terminal (with your json-array-formatted file)
    ```cgo
    $ mongoimport --uri "mongodb://root:root@localhost:27017/hanako" --collection anime --file formatted-anime-offline-database.json --jsonArray
    ```

## Setup
1. Create a ``.env`` file on the folder
2. Setup the dev environment there, format
```dotenv
mongo_url=mongodb_url
mongo_dbname=mongo_db_database_name
mongo_username=root
mongo_pwd=root
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