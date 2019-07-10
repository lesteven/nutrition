#!/bin/bash

read -p "Enter val: " val


if [ "$val" = "" ]; then
    echo "no value given"
    sudo docker run -d -p 8080:8080 --name nutserver lesteven/nutserver
else
    echo "value is $val"
    sudo docker run -e ADDRESS="$val" -d -p 8080:8080 --name nutserver lesteven/nutserver
fi
