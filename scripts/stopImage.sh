#!/bin/bash

sudo docker container stop nutserver
sudo docker container rm nutserver
sudo docker container prune
