#!/bin/sh

echo "create rest service on k8s cluster...."

exist=`kubectl get svc --namespace=unit-test | grep sys | grep -v grep`
if [ "$exist" = "" ]
then
  echo "sys service not exist! NOT need delete "
else
  kubectl delete svc sys --namespace=unit-test
fi

kubectl create -f ./Service.json
echo "sys service done"

echo "create sys ReplicationController on k8s cluster...."
exist=`kubectl get ReplicationController --namespace=unit-test | grep sys |grep -v grep`
if [ "$exist" = "" ]
then
  echo "sys ReplicationController not exist! NOT need delete"
else
  kubectl delete ReplicationController sys-rc --namespace=unit-test
fi

kubectl create -f ./ReplicationController.json
echo "sys ReplicationController done"

