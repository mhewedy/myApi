#!/usr/bin/env bash

## generate swagger docs
rm -rf docs
swag init -g cmd/api/main.go


# gin is a Live reload utility for Go web servers
# https://github.com/codegangsta/gin

# gin runs on port 3000 (by default) and forward traffic to port 8080
gin -i -a 8080 -b .gin-bin --all -d cmd/api run main.go

