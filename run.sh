#! /bin/bash

./stop.sh
./rmi.sh

clear
reset

BUILD=local
APP_NAME=server-ping

project_path="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
echo "project_path = ${project_path}"

docker-compose -f docker/docker-compose.local.yml up --build