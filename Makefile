NAME=godev
VERSION=1.0

MAKEFILE_DIR := $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

build:
	docker build -t godev .

start:
	docker run -it -v $(MAKEFILE_DIR):/root/work godev

drop:x
	docker image 