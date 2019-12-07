# Project Hanako [![Build Status](https://travis-ci.com/luqmansen/Hanako.svg?branch=master)](https://travis-ci.com/luqmansen/Hanako)

![](docs/assets/hanako-small.png)


### What is This ?
For now, this is a RESTful Webservice API for an anime database

## Demo
(V1 use postgres)<br>
- https://hanako.luqmansen.me/api/v1/anime/all
- https://hanako.luqmansen.me/api/v1/anime/search/q?title=kimi

(V2 use mongodb)<br>
- https://hanako.luqmansen.me/api/v2/anime/all
- https://hanako.luqmansen.me/api/v2/anime/search/q?title=kimi


## Prequisites

#### Tools
* [Go](https://golang.org/)
* [Postgresql](https://www.postgresql.org)
* [MongoDB](mongodb.com)
* [Compass ](https://www.mongodb.com/products/compass) (Optional, but easier for development)

#### Packages Depencency( For Development)
* gorilla/mux
* jinzhu/gorm
* dgrijalva/jwt-go
* joho/godotenv
* tools/godep
* gopkg.in/mgo.v2
* gopkg.in/mgo.v2/bson

## Instalation
```bash
git clone https://github.com/luqmansen/Hanako.git
```
or
```bash
go get https://github.com/luqmansen/Hanako
``` 

### Setup
1. Create a ``.env`` file on the folder
2. Setup the dev environtment there, format
```dotenv
db_name = your_db_name
db_pass = your_db_name
db_user = db_username
db_type = postgres
db_host = localhost
db_port = 5434 (Postgres port, left it for default)
token_password = jwt_Secret_Token
mongo_url = mongodb_url
mongo_dbname = mongo_db_database_name
```

## Usage
Run the server
```markdown
cd Hanako
go build main.go
./main
```
Available endpoints
```
/api/v1/user/new
/api/v1/user/login
/api/v1/anime/all
/api/v1/anime/{id}
/api/v1/anime/search
```

#### Query keyword 
```
/api/v1/anime/search/q?title=anime_to_find
/api/v1/anime/all?show=1000
``` 

Other stuff still on development, fork this repo and help me :D

## Deployment

- [Heroku](https://github.com/luqmansen/Hanako/wiki/Deployment)
 
## Roadmap and Progress 
Read the roadmap on the wiki [here](https://github.com/luqmansen/Hanako/wiki/Roadmap)

## Contributing
For now, just fork it and create a pull request, there will be a contribution guidelines soon!

## Authors

<a href="https://github.com/luqmansen/Hanako/graphs/contributors">
  <img src="https://contributors-img.firebaseapp.com/image?repo=luqmansen/Hanako" />
</a>


## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/luqmansen/hanako/blob/master/docs/LICENSE.md) file for details
