#!/bin/sh
src="main.go"
env=".env.dev"

# Start pgAdmin Container
printf "\nRunning pgAdmin from docker\nRemeber to kill the container manually\n"
docker run --name pgAdmin --rm -p 5555:80 --env-file $env -d dpage/pgadmin4:latest
printf "pgAdmin container is running at localhost:5555\n"

# Start trips_pg PostgreSQL container
printf "\nRunning trips_pg from docker\nRemeber to kill the container manually\n"
docker run --name trips_pg --rm -p 5432:5432 --env-file $env -d postgres:13
host=$(docker inspect --format "{{.NetworkSettings.IPAddress}}" trips_pg)
printf "trips_pg container is running at $host:5432\n"

# Start the server
printf "\nStart running: server ...\n"
# Set all ENV vars for the server to run
export $(grep -v '^#' $env | xargs) && export POSTGRES_HOST="$host" && go run $src
# This should unset all the ENV vars, just in case.
unset $(grep -v '^#' $env | sed -E 's/(.*)=.*/\1/' | xargs)
unset POSTGRES_HOST
printf "\nStopped running: server\n\n"