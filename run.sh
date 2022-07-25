#!/bin/bash

docker build -f Dockerfile -t flagship-cli .
docker run --rm -it --name cli -t flagship-cli