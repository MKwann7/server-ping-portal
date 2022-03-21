#! /bin/bash

project_path="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
echo "project_path = ${project_path}"

docker-compose -f docker/docker-compose.local.yml down

sudo chmod -R 777 "${project_path}/docker/mysql-data"
sudo rm -rf "${project_path}/docker/mysql-data/*"

docker container rm -f $(docker container ls -a -q)
docker rmi -f $(docker images -a -q)
docker volume rm -f $(docker volume ls -q)
docker system prune --volumes

echo "excell application removed"