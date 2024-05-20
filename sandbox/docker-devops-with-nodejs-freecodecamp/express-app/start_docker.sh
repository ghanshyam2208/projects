#!/bin/bash

# possible commands
# ./start_docker.sh down
# ./start_docker.sh build 1
# ./start_docker.sh volume 1
# ./start_docker.sh rebuild 2
# ./start_docker.sh rebuild 1

if [ "$1" = "build" ]; then
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml up --build -d --scale my-node-app=$2
    sleep 2
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml  logs -f my-node-app
elif [ "$1" = "down" ]; then
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml down
elif [ "$1" = "volume" ]; then
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -V -d --scale my-node-app=$2
    sleep 2
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml  logs -f my-node-app
elif [ "$1" = "rebuild" ]; then
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml down
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml up --build -d --scale my-node-app=$2
    sleep 2
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml  logs -f my-node-app
else
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -V -d 
    sleep 2
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml  logs -f my-node-app
fi