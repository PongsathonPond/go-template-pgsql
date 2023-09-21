#!bin/bash

NAME="idev-cms-service"
docker build -f ./deployment/Dockerfile-sit -t $NAME:sit .