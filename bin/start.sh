#!/bin/sh
cd /uip/bin
./uip -dbip=${SSDB_IP} -dbport=${SSDB_PORT} -servicePort=${SERVICE_PORT} -p=${SVC_PORT} -etcdport=${ETCD_PORT} -etcdip=${ETCD_IP} -logip=${LOG_IP} -logport=${LOG_PORT} -pseIp=${PSE_IP} -psePort=${PSE_PORT} -connMinutes=${CONN_MINUTES}

