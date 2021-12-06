#!/bin/bash
bold=$(tput bold)
boldend=$(tput sgr0)
echo "Starting Go Project in ${bold}DEV Mode${boldend}...."
exec go run main.go