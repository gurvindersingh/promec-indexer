#!/bin/bash
TAG=$1

docker build -t quay.io/$TRAVIS_REPO_SLUG:$TAG .;
docker login -u="$DOCKER_USERNAME" -p="$DOCKER_PASSWORD" quay.io;
docker push quay.io/$TRAVIS_REPO_SLUG:$TAG;
