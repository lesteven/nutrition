#!/bin/bash

read -p "Enter val: " val


if [ "$val" = "" ]; then
    echo "no value given"
    sudo docker run --net=host -d -p 8080:8080 --name nutserver lesteven/nutserver:dev
else
    echo "value is $val"
    sudo docker run --net=host  -e ADDRESS="$val" -d -p 8080:8080 --name nutserver lesteven/nutserver:dev
fi
