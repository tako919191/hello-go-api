#!/bin/sh

if [ $1 = 'minikube' ]; then
    eval $(minikube docker-env)
    echo 'Build on Minikube VM...'
else
    echo 'Build on Local Machine...'
fi

docker build -t tako919191/hello-go-api:latest .

echo 'Done!'
