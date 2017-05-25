#!/bin/bash

#run fabric vp2
cd /home/a/fabric/vp2
echo 'a11111'|sudo -S nohup docker-compose -f server_vp2.yml up &
echo 'run fabric vp2'