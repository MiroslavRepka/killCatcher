#!/bin/bash
if [ $# -eq 0 ]
  then
    echo "Please, choose which of the test directories you wish to test with"
    echo "Options:"
    echo "  test-continue-main => test where main logic will continue after SIGTERM"
    echo "  test-end-main => test where main logic will end after SIGTERM"
    exit
fi

echo "---== Testing option: $1 ==---"
TEST_PATH=$1
export MINIKUBE_IN_STYLE=0

echo "---== Create minikube cluster ==---"
minikube start

echo "---== Build docker image for minikube ==---"
eval $(minikube docker-env)
docker build -f Dockerfile -t kill-catcher-test:v1 $TEST_PATH

echo "---== Deploy test pod ==---"
kubectl apply -f pod.yaml

echo "---== Wait for pod to be ready ==---"
kubectl wait pods -l app=test-app --for condition=Ready --timeout=300s

echo "---== Cluster and pod are ready! ==---"
echo "---== Wait until everything is initilised ==---"
sleep 10

echo "---== Deleting test pod ==---"
kubectl delete -f pod.yaml&

echo "---== Saving pod logs ==---"
LOGS=""
while true;
do
  LOGS=$(kubectl logs test-app) 
  if [[ "$?" -ne 0 ]]; then 
    break
  fi
  echo "$LOGS" > test-app.log
  sleep 1
done

echo "---== Deleting cluster ==---"
minikube delete

echo "To see logs, please see the test-app.log file"
