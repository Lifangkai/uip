#!/bin/sh

echo "create uip rest service on k8s cluster...."

exist=`kubectl get svc --namespace=unit-test | grep uip | grep -v grep`
if [ "$exist" = "" ]
then
  echo "uip service not exist! NOT need delete "
else
  kubectl delete svc uip --namespace=unit-test
fi

kubectl create -f ./Service.json
echo "uip service done"

echo "create uip ReplicationController on k8s cluster...."
exist=`kubectl get ReplicationController --namespace=unit-test | grep uip-rc |grep -v grep`
if [ "$exist" = "" ]
then
  echo "uip ReplicationController not exist! NOT need delete"
else
  kubectl delete ReplicationController uip-rc --namespace=unit-test
fi

kubectl create -f ./ReplicationController.json
echo "uip ReplicationController done"

