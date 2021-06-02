# myApi

Opinionated modular template for starting an golang api using:
* [echo](https://echo.labstack.com/) web framework.
* [gorm](http://gorm.io) to access DB.
* Postgres as database.
* [go-migrate](https://github.com/golang-migrate/migrate) for db schema migration.
* [swaggo](https://github.com/swaggo/swag) to generate swagger.
* [gin](https://github.com/codegangsta/gin) to live reload the go api.


It contains:
* simple startup app with simple registration & login APIs that provides JWT-based authentication. 
* health check endpoint.
* Multistage Dockerfile
* Gitlab CI script.

### Steps:

1. run `./rename` to rename the application. e.g. `./rename todoApi`

2. run `./etc/create-db` to create docker-based postgres database

3. run `go get -d ./...` to install all dependencies

4. run `./start` to start the app   
it requires the following golang utilities:
```
go get -u github.com/codegangsta/gin			# for gin the live reload utility
go get -u github.com/swaggo/swag/cmd/swag		# for swag the swagger generation utility

```
5. access the api using: http://localhost:8080/swagger/index.html
