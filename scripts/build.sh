#!/bin/sh
output="build/server"
src="cmd/server/main.go"

printf "\nBuilding: server ...\n"
if [ $# -eq 0 ]
then
  go build -o $output $src
elif [ "$1" = "no_cgo" ]
then
  CGO_ENABLED=0 go build -o $output $src
else
  printf "Provide a valid argument ('no_cgo' for disableing CGO compiler)\n"
  exit 1
fi
printf "\nBuilt: server size:"
ls -lah $output | awk '{print $5}'
printf "\nDone building: server\n"
printf "\nNote: Remember to load the env vars before running the binary\n\n"