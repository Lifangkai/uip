#!/bin/sh
cd /uip/bin
./uip -dbip=${SSDB_IP} -dbport=${SSDB_PORT} \
-servicePort=${SERVICE_PORT} -p=${SVC_PORT} \
-etcdport=${ETCD_PORT} -etcdip=${ETCD_IP} \
-logip=${LOG_IP} -logport=${LOG_PORT} \
-pseIp=${PSE_IP} -psePort=${PSE_PORT} \
-connMinutes=${CONN_MINUTES} \
-ext=${EXT}


#nohup ./uip -dbip=172.16.0.13 -dbport=7000 \
#-servicePort=7001 -p=7001 \
#-etcdport=2379 -etcdip=172.16.0.11 \
#-logip=172.16.0.13 -logport=7702 \
#-pseIp=172.16.0.13 -psePort=7502 \
#-connMinutes=5 \
#-ext="{\"svcaddr\":\"cl\", \"SNSIP\":\"172.16.0.13\",\"SNSPORT\":\"7501\"}" \
#> ../logs/nohup.log 2>&1 &

