#!/usr/bin/env bash

protoc -I=. --gofast_out=../client *.proto
protoc -I=. --gograinv2_out=../client *.proto
