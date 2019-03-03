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

