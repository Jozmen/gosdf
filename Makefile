get_sdf:
	./tools/get_sdf.sh

generate:
	./tools/generate.sh

build: build_convert

build_convert: generate
	go build -o ./gosdf-convert ./cmd/gosdf-convert/main.go

install: generate
	go install -i ./cmd/gosdf-convert