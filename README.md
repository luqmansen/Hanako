# Project Hanako

![](docs/assets/hanako-small.png)

Formerly, this was [ProjectAmbis](https://luqmansen.github.io/project-ambis-started/), an abstract project that i want to make but it was too abstract that i haven't any idea where to start, now its here

### What is This ?
For now, this is a RESTful Webservice API for an anime database


### Prequisites

##### Tools
* [Go](https://golang.org/)
* [Postgresql](https://www.postgresql.org)

##### Packages ( For Development)
* gorilla/mux
* jinzhu/gorm
* dgrijalva/jwt-go
* joho/godotenv

### Instalation
``
git clone https://github.com/luqmansen/Hanako.git
cd Hanako
go run main.go
``

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
token_password = jwt_Secret_Token;

```

### Usage
Access this from on wherever your deployment is
```
/api/v1/user/new
/api/v1/user/login
/api/v1/anime/all
/api/v1/anime/{id}
/api/v1/anime/search

```

##### Query keyword 
```/api/v1/anime/search?title=anime_to_find``` 


### Roadmap and Progress 
Planned to create a wiki, for now, you can read it on my [blog](https://luqmansen.github.io) and
I try to implementing another [techstack](https://github.com/luqmansen/hanako/blob/master/docs/Techstack.md)


### Contributing
For now, just fork it and create a pull request, i haven't made any contribution guidelines right now

## Authors

* **Luqmansen** - [luqmansen](https://github.com/luqmansen)

See also the list of [contributors](https://github.com/luqmansen/hanako/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/luqmansen/hanako/blob/master/docs/LICENSE.md) file for details
