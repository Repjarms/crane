#!/bin/bash

# pull from docker
echo "Pulling from docker"

REPO_NAME=$1
echo $REPO_NAME


# get container id
CONTAINER_ID=$(docker ps -a -q --filter ancestor=$REPO_NAME)
echo "Found existing container at $CONTAINER_ID"

# stop existing container
echo "stopping existing container $CONTAINER_ID"
echo "docker stop $CONTAINER_ID" | bash

# remove existing container
echo "removing existing container $CONTAINER_ID"
echo "docker rm $CONTAINER_ID" | bash

# remove existing image
echo "docker rmi $1" | bash

# pull latest image, defaults to latest tag
echo "docker pull $1" | bash

# run new container using latest image
echo "docker run -d -p 8012:8012 $1" | bash

# proxy the ip address to something that is reachable
