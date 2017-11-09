#!/bin/bash

rm -rf app
go build app.go
echo "build suc"
chmod +x app
echo "=============="
./app
