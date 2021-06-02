# myApi

### What is: 
Opinionated modular template for starting  Golang API. It uses:
* [echo](https://echo.labstack.com/) web framework.
* [gorm](https://gorm.io/) ORM library.
* [validator](https://github.com/go-playground/validator) to validate DTOs.
* [go-conf](https://github.com/mhewedy/go-conf) to read external configurations.
* Postgres as database.
* [go-migrate](https://github.com/golang-migrate/migrate) for db schema migration.
* [swaggo](https://github.com/swaggo/swag) to generate swagger.
* [gin](https://github.com/codegangsta/gin) to live reload the go api.


It contains:
* simple startup app with simple registration & login APIs that provides JWT-based authentication. 
* health check endpoint.
* Multistage Dockerfile
* Gitlab CI script.
* Database schema migration files found under `db/migrations`

### How it differs from Xyz:
Unlike other templates or tools, myApi is adapting the idea of modular monolithic, in which you can use  `pkg/<module_name>` to host different modules, (e.g. `pkg/user` and `pkg/health`).    

The code in the different modules can access shared services via the `pkg/commons` modules. e.g `commons.DB` to access gorm DB object and `commons.Validate` to access Validate object. (You can put shared services here too, such as `Redis` Client)

### Module Structure:
each module (e.g. `pkg/user`, `pkg/main_business_module`) contains:
1. `model.go` to contains all model structs and related db-access methods in this module.
2. `dto.go` contains DTOs.
3. one or more `<some_function>.go` files (e.g. `login.go`, `register.go`) which contains the [echo](https://echo.labstack.com/) callback function that it registered at `cmd/api/routes/route.go`

## How to use it:

1. Clone the repo or download the [zip file](https://github.com/mhewedy/myApi/archive/refs/heads/master.zip).

2. run `./rename` to rename the application. e.g. `./rename todoApi`

3. run `./etc/create-db` to create docker-based Postgres database

4. run `go get -d ./...` to install all dependencies

5. run `./start` to start the app   
it requires the following Golang utilities:
```
go get -u github.com/codegangsta/gin			# for gin the live reload utility
go get -u github.com/swaggo/swag/cmd/swag		# for swag the swagger generation utility

```
6. access the API using: http://localhost:8080/swagger/index.html
