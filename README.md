# Project Hanako

![](docs/assets/hanako-small.png)

Formerly, this was [ProjectAmbis](https://luqmansen.github.io/project-ambis-started/), an abstract project that i want to make but it was too abstract that i haven't any idea where to start, now its here

### What is This ?
For now, this is a RESTful Webservice API for an anime database

## Demo
* Host : https://hanako.luqmansen.me/ <br>
Index will return 404, user other endpoints, eg:
    - https://hanako.luqmansen.me/api/v1/anime/all
    - https://hanako.luqmansen.me/api/v1/anime/search/q?title=kimi

## Prequisites

#### Tools
* [Go](https://golang.org/)
* [Postgresql](https://www.postgresql.org)

#### Packages Depencency( For Development)
* gorilla/mux
* jinzhu/gorm
* dgrijalva/jwt-go
* joho/godotenv
* tools/godep

## Instalation
```bash
git clone https://github.com/luqmansen/Hanako.git
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
token_password = jwt_Secret_Token;

```

## Usage
Run the server
```markdown
cd Hanako
go run main.go
```
Access this from on wherever your deployment is, in this example, on localhost
```
localhost:8080/api/v1/user/new
localhost:8080/api/v1/user/login
localhost:8080/api/v1/anime/all
localhost:8080/api/v1/anime/{id}
localhost:8080/api/v1/anime/search

```

#### Query keyword 
```
/api/v1/anime/search/q?title=anime_to_find
``` 

Other stuff still on development, fork this repo and help me :D

## Deployment

### Heroku 
___
##### Prequites
1. Heroku-cli installed on your machine 
    ```
    curl https://cli-assets.heroku.com/install-ubuntu.sh | sh
    ```
2. Create your app name
    ```
    heroku create <your_app_name>
    ```
3. Heroku will do its job, then push it
    ```
    git push heroku master
    ```
4. Add Postgresql addons on heroku
    ```
    heroku addons:create heroku-postgresql:hobby-dev
    ```
   
5. Now open heroku on the browser, go to setting -> Reveal Config vars
![](docs/assets/heroku-config-vars.png) <br>
Look at the ```DATABASE_URL```
The format is 
    ```
    postgres://username:password@host/dbName
    ```

6. Add the config key and value <br>
 ![](docs/assets/create-config.png)
___
##### Importing the database
In case you working with local database and want to import it to the remote

1. Dump your local database
    ```
    $ PGPASSWORD=yourpassword pg_dump -Fc --no-acl --no-owner -h localhost -U <your_db_username>  <your_dbname> > database.dmp
    ```
2. Restore it to remote database
    ``` 
    pg_restore --verbose --clean --no-owner --no-acl -h <hostname> -U <username> -d <database_name> -p 5432 database.dump
    ```
3. Enter the password on prompted line
4. Cool :D


## Roadmap and Progress 
Planned to create a wiki, for now, you can read it on my [blog](https://luqmansen.github.io) for the progress that i made while trying to implementing another [techstack](https://github.com/luqmansen/hanako/blob/master/docs/Techstack.md), you can help here :D


## Contributing
For now, just fork it and create a pull request, there will be a contribution guidelines soon!

## Authors

* **Luqmansen** - [luqmansen](https://github.com/luqmansen)

See also the list of [contributors](https://github.com/luqmansen/hanako/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/luqmansen/hanako/blob/master/docs/LICENSE.md) file for details
