#!/bin/sh

IMAGE_NAME=vspeach/crud-service
IMAGE_VERSION=4.0

docker build -t $IMAGE_NAME:$IMAGE_VERSION .
docker push $IMAGE_NAME:$IMAGE_VERSION