#!/bin/bash


cd "$(dirname "$0")"
cd ..
sudo docker build -t lesteven/nutserver .
