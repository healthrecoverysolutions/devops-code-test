#!/bin/bash

cp requirements.txt docker/python-v3

sudo docker build -t build-python-v3 docker/python-v3

sudo docker run -it --rm --name python-version-two -v "$PWD":$(pwd) -w $(pwd) build-python-v3 python main.py