#!/bin/bash
docker rm -f tally
docker run -itd --network bridge \
 -p 6666:8888 \
 -v /usr/share/zoneinfo/Asia/Shanghai:/etc/timezone:ro \
 -v /etc/localtime:/etc/localtime:ro \
 -v /etc/hosts:/etc/hosts:ro \
 -v `pwd`/server:/opt/work \
 -w /opt/work \
 --restart always --name tally \
 ubuntu:18.04 ./server -c config.docker.yaml
