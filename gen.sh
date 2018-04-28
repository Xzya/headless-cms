#!/usr/bin/env bash

mkdir -p gen
swagger generate server -t gen -f ./swagger/swagger.yml --exclude-main -A cms