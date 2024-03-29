#! /bin/bash

project_path="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
echo "Uninstalling project at path: ${project_path}"

docker-compose -f docker/docker-compose.local.yml down --volumes

echo "Removing database: ${project_path}/docker/postgres-data/*"
sudo chmod -R 777 docker/postgres-data
sudo rm -rf docker/postgres-data/*

docker rm serverping-app
docker rmi serverping_app
echo "y" | docker system prune --volumes -a

echo "server ping application removed!"