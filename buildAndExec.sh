#!/bin/sh
cd ~/go/src/github.com/michalnov/SovyGo/
#sh setup_web_content.sh
#cd ..
go build -o build/server bin/main.go
./build/server
wait