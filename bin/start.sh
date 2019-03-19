#!/bin/sh

nohup ./uip -dbip=172.16.0.13 -dbport=7000 \
-servicePort=7011 -p=7011 \
-etcdport=2379 -etcdip=172.16.0.11 \
-logip=172.16.0.13 -logport=7702 \
-pseIp=172.16.0.13 -psePort=7502 \
-connMinutes=5 \
-ext="{\"svcaddr\":\"cl\", \"SNSIP\":\"172.16.0.13\",\"SNSPORT\":\"7501\"}" \
> ../logs/nohup.log 2>&1 &

