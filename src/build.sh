#!/usr/bin/env bash

rm -rf /go/src/github.com/ralali

mkdir vendor
cp -rf /go/src/* vendor/

mkdir -p /go/src/github.com/ralali/

ln -s /my_app /go/src/github.com/ralali/rl-ms-boilerplate-go

cd /go/src/github.com/ralali/rl-ms-boilerplate-go
dep ensure -v

cd /my_app

cp -rf vendor/* /go/src
rm -rf vendor

refresh init -c config.yml
refresh run -c config.yml