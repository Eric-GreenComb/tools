#!/bin/bash

#run rsa server
# rsaserver='cd /home/eric/go/src/github.com/CebEcloudTime/rsaserver ; nohup ./main &'
# echo ${rsaserver}|awk '{run=$0;system(run)}'
cd /home/a/rsaserver
nohup ./main &
echo 'run rsa server'

#run fabric vp0
cd /home/a/fabric/vp0
echo 'a11111'|sudo -S nohup docker-compose -f server_vp0.yml up &
echo 'run fabric vp0'

#run charityapp
cd /home/a/app
nohup java -jar charityapp-1.0.jar &
