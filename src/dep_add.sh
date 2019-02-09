#!/usr/bin/env bash

mkdir -p /go/src/github.com/ralali;
rm -rf /go/src/github.com/ralali/rl-ms-boilerplate-go;
ln -s /my_app /go/src/github.com/ralali/rl-ms-boilerplate-go;
cd /go/src/github.com/ralali/rl-ms-boilerplate-go;
dep ensure -add $1;
go get $1;
rm -rf vendor