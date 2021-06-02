# myApi

Opinionated template for starting an golang api using [echo](https://echo.labstack.com/), [gorm](http://gorm.io/index.html), Postgres and [swaggo](https://github.com/swaggo/swag).

It contains simple startup app with simple registration & login that provides JWT-based authentication. 

### prerequisite:

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
