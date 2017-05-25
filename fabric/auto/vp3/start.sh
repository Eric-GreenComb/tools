#!/bin/bash

#run fabric vp3
cd /home/a/fabric/vp3
echo 'a11111'|sudo -S nohup docker-compose -f server_vp3.yml up &
echo 'run fabric vp3'