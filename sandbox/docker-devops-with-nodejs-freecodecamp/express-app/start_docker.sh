#!/bin/bash

if [ "$1" = "build" ]; then
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml up --build -d
    sleep 2
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml  logs -f my-node-app
elif [ "$1" = "down" ]; then
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml down
elif [ "$1" = "volume" ]; then
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -V -d 
    sleep 2
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml  logs -f my-node-app
elif [ "$1" = "rebuild" ]; then
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml down
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml up --build -d
    sleep 2
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml  logs -f my-node-app
else
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -V -d 
    sleep 2
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml  logs -f my-node-app
fi