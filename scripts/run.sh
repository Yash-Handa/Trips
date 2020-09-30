#!/bin/sh
src="cmd/server/main.go"
env=".env"

printf "\nStart running: server ...\n"
# Set all ENV vars for the server to run
export $(grep -v '^#' $env | xargs) && go run $src
# This should unset all the ENV vars, just in case.
unset $(grep -v '^#' $env | sed -E 's/(.*)=.*/\1/' | xargs)
printf "\nStopped running: server\n\n"