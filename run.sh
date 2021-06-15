#!/usr/bin/env bash

docker build -t app .
docker run -i -t -p 8080:8080 app
