#!/bin/sh

SERVICE_VERSION=1.0

docker run --rm -it -v $PWD:/opt/echoserver golang:1.13-alpine sh -c "cd /opt/echoserver/src && go build -o ../bin/echoserver ./module/echoserver"

/bin/cp -f $PWD/bin/echoserver $PWD/docker/

docker build --no-cache -t echoserver:${SERVICE_VERSION} $PWD/docker/
