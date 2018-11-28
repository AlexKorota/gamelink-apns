#!/bin/sh
PROJECT=gamelink-apns
RELEASE=1.0
COMMIT=`git rev-parse --short HEAD`
BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S')
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-X ${PROJECT}/version.Release=${RELEASE} -X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}"
docker build -t mrcarrot/gamelink-apns .
docker push mrcarrot/gamelink-apns