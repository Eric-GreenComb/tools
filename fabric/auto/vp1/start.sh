#!/bin/bash

#run fabric vp1
cd /home/a/deploy/vp1
echo 'a11111'|sudo -S nohup docker-compose -f server_vp1.yml up &
echo 'run fabric vp1'