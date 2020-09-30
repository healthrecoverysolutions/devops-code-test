#!/bin/bash

cp requirements.txt docker/python-v2

sudo docker build -t build-python-v2 docker/python-v2

sudo docker run -it --rm --name python-version-two -v "$PWD":$(pwd) -w $(pwd) build-python-v2 python main.py