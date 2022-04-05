#!/usr/bin/env bash


###
## install dependencies
###

###
## prepare GO App
###
go get -u
go mod tidy
go mod download

###
## build GO App
###
