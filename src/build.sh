#!/usr/bin/env bash

rm -rf /go/src/github.com/ralali
mkdir -p /go/src/github.com/ralali/

ln -s /my_app /go/src/github.com/ralali/rl-ms-boilerplate-go

cd /go/src/github.com/ralali/rl-ms-boilerplate-go
dep ensure -v -vendor-only

cd /my_app
rm -rf /go/src/github.com/ralali

cp -rf vendor/* /go/src
rm -rf vendor

refresh init
refresh run