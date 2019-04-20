[![Build Status](https://travis-ci.org/tpageforfunzies/boulder.svg?branch=master)](https://travis-ci.org/tpageforfunzies/boulder)
[![GolangCI](https://golangci.com/badges/github.com/golangci/golangci-lint.svg)](https://golangci.com)

# boulder 
third iteration on the bouldertracker idea, this time will be a go service with separate web and mobile clients, this repo is for the Go web service and will be using the gin framework, probably with gorm for orm eventually and a database of some sort.  probably postgres

# to try it out
you'll need: 
  * Go >1.11
  * Dep
  * Make

clone the repo into your go workspace, get into github.com/tpageforfunzies/boulder/ and run `make setup` and then `make basic` and it'll install dependencies and spin up the local web server, visit `localhost:1337/v1/` and see the json response

if you dont want to set up a go workspace and install go 
you can also run it in a docker container (this will move source and .env into container and build there so don't have to build locally if you dont want to), something like this'll do it
<br>
build image 
<br>
`docker build -t boulderlinux .`
<br>
create the container
<br>
`docker create -it --name boulderapp -p 80:80 boulderlinux:latest`
<br>
start the container
<br>
`docker start boulderapp`
<br>
of course now there's an even easier `dockerdeploy.sh` that does all that for you

# Back End To Do:
 * API
 	* Add pagination to V1 api
 * Deployment
 	* Improve docker scripts (clean up better)
 	* Integrate Nginx RP repo to this repo and deployment
 * Feature
 	* User Friends, Friend Feed
 	* Locations
 	* Images
 * Documentation
 	* Document SSL/Deployment
 	* Update swagger spec for HTTPS/WWW
 * Testing
 	* Start
