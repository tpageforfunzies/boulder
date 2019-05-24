#!/bin/bash
# stop and delete old container
# this could be fancier and check a docker ps for the container
docker stop boulderapp && docker rm boulderapp
# build new image
# this could read a version env var and tag itself
docker build -t boulderlinux:v1 .
# create new container
# this could be a template file
docker create -it --name boulderapp -p 8080:80 boulderlinux:v1
# start that puppy up
docker start boulderapp
