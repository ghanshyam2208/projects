#!/bin/bash
if [ "$1" = "db" ]; then
    docker-compose -f ./docker-compose-db.yml up --build -d postgres-service pgadmin-service

elif [ "$1" = "build" ]; then
    docker-compose -f ./docker-compose.yml up --build -d
    sleep 2
    docker-compose -f docker-compose.yml logs -f ms-users-mgmt ms-auth 
elif [ "$1" = "down" ]; then
    docker-compose -f ./docker-compose.yml down
elif [ "$1" = "logs" ]; then
    docker-compose -f docker-compose.yml logs -f ms-users-mgmt ms-auth 
elif [ "$1" = "dbdown" ]; then
    docker-compose -f ./docker-compose-db.yml down
else
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -V -d
    sleep 2
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml  logs -f my-node-app
fi