#!/usr/bin/env bash

if [ -z "$1" ] || [ -z "$2" ]
then
    echo "usage: image-tagpush <tag> <user>"
fi

TAG="${1:-latest}"
HUBUSER="${2:-ubombar}"

echo "Running with arguments tag='$TAG' user='$HUBUSER'"

for name in ./build/images/*/
do 
    bname=$(basename $name)
    docker tag $bname:latest $HUBUSER/$bname:$TAG
    docker push $HUBUSER/$bname:$TAG
done 