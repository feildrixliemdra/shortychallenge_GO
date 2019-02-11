#!/usr/bin/env bash

rm -rf .vendor-new
rm -rf /go/src/github.com/ralali

mkdir -p vendor
cp -rf /go/src/* vendor/

mkdir -p /go/src/github.com/ralali/
touch /go/src/github.com/.gitignore

ln -s /my_app /go/src/github.com/ralali/rl-ms-boilerplate-go

cd /go/src/github.com/ralali/rl-ms-boilerplate-go
dep ensure -v

cd /my_app

cp -rf vendor/* /go/src
rm -rf vendor

swag init