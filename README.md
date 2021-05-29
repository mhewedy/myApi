# myApi

### prerequisite:

1. run `./rename.sh` to rename the application. e.g. `./rename.sh todoApi`

2. run `./create-db.sh` to create docker-based postgres database

3. run `go get -d ./...` to install all dependencies

4. run `./run.sh` to run the app   
it requires the following golang utilities:
```
go get -u github.com/codegangsta/gin			# for gin the live reload utility
go get -u github.com/swaggo/swag/cmd/swag		# for swag the swagger generation utility

```

access the api using: http://localhost:8080/swagger/index.html
