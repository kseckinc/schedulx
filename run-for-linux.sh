#!/usr/bin/env bash
# deploy mysql
docker run -d --name schedulx_db -e MYSQL_ROOT_PASSWORD=mtQ8chN2 -e MYSQL_DATABASE=schedulx -e MYSQL_USER=gf -e MYSQL_PASSWORD=db@galaxy-future.com -p 3316:3306 -v $(pwd)/init/mysql:/docker-entrypoint-initdb.d yobasystems/alpine-mariadb:10.5.11
# deploy schedulx
docker run -d --name schedulx_api --network host -v $(pwd)/register/conf/config.yml:/home/schedulx/register/conf/config.yml galaxyfuture/schedulx-api bin/wait-for-schedulx.sh