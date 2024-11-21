#!/bin/sh

go build -o bin/migration cmd/migration/main.go
./bin/migration

go build -o bin/main cmd/api/main.go
./bin/main