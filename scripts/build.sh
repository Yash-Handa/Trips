#!/bin/sh
output="build/server"
src="cmd/server/main.go"

printf "\nBuilding: server ...\n"
go build -o $output $src
printf "\nBuilt: server size:"
ls -lah $output | awk '{print $5}'
printf "\nDone building: server\n\n"