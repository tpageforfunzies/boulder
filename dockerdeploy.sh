#!/bin/bash
# update repo
git pull
# compile mac and linux
make build
# stop and delete old container
docker stop boulderapp && docker rm boulderapp
# build new image
docker build -t boulderlinux .
# create new container
docker create -it --name boulderapp -p 80:80 boulderlinux:latest
# start that puppy up
docker start boulderapp