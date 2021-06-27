#!/usr/bin/env bash

# execute the script after successful build and image pushed to the a docker registry (gitlab)
# to redeploy the new build

docker-compose pull
docker-compose up --force-recreate -d
