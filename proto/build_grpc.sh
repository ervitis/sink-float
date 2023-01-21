#!/usr/bin/env bash

docker run \
	--rm -v "${PWD}":/defs namely/protoc-all:1.51_1 \
	-f ./*.proto \
	-l go \
	--go_out=plugins=grpc:/gen/grpc