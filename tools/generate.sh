#!/bin/bash

go run cmd/gosdf-generate/main.go -s ./sdformat/sdf/1.5
go fmt github.com/Jozmen/gosdf/pkg/sdf