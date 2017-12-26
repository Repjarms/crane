#!/bin/bash

# pull from docker
echo "Pulling from docker"

REPO_NAME=$1
echo $REPO_NAME


# loop through all containers spawned from the related image and stop/remove
# them

echo "Finding all containers running current image"

for con in $(docker ps -a -q --filter ancestor=$REPO_NAME); do
  # stop current container
  echo "stopping existing container $con"
  echo "docker stop $con" | bash

  # remove current container
  echo "removing existing container $con"
  echo "docker rm $con" | bash
done

# remove existing image
#echo "docker rmi $1" | bash

# pull latest image, defaults to latest tag
echo "docker pull $1" | bash

# get exposed port from the docker config
APPLICATION_PORT=$(docker inspect --format='{{.Config.ExposedPorts}}' $1 | tr -dc 0-9)
echo "application exposed port $APPLICATION_PORT"

# run new container using latest image
echo "docker run -d -p $APPLICATION_PORT:$APPLICATION_PORT $1" | bash
echo "application running latest image of $1 on port $APPLICATION_PORT"

