#!/usr/bin/env bash


###
## install dependencies
###

###
## prepare GO App
###
go mod tidy
go mod download

###
## build GO App
###
