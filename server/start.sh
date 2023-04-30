#!/bin/bash
docker run -itd --network host \
 -p 8888:8888 \
 -v /usr/share/zoneinfo/Asia/Shanghai:/etc/timezone:ro \
 -v /etc/localtime:/etc/localtime:ro \
 -v /etc/hosts:/etc/hosts:ro \
 -v `pwd`:/opt/work \
 -w /opt/work \
 --restart always --name mini_app \
 ubuntu:18.04 ./main

