#! /bin/bash

docker build . -t npm-analysis
docker run --rm --name npm-analysis -p 8983:8983 npm-analysis