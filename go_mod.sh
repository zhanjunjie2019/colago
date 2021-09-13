#!/usr/bin/env bash

cd common
go mod tidy

cd ..
cd samples

cd auth-domain
go mod tidy

cd ..
cd auth-client
go mod tidy

cd ..
cd user-domain
go mod tidy

cd ..
cd user-client
go mod tidy

cd ..
cd test
go mod tidy
