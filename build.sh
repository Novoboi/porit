#!/bin/bash

# Defining Useful Variables For Echo
bold=$(tput bold)
boldend=$(tput sgr0)

#================================[ Executing Commands ]================================#
echo "Building Project to ${bold}web-service-gin${boldend} file in ${bold}./${boldend}"
exec go build
#======================================================================================#
