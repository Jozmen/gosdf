# gosdf

gosdf is a tool for creating sdf models using yaml.
As input it takes xml schema from original sdformat repository (https://bitbucket.org/osrf/sdformat) and generates code for creating sdf objects from yaml files.

## Setup

1. Install GoLang https://golang.org/doc/install
2. Get gosdf:
```
go get github.com/Jozmen/gosdf
```
3. Enter project directory:
```
cd $GOPATH/src/github.com/Jozmen/gosdf
```
4. get the sdformat:
```
make get_sdf
```

## Build

```
make build
```

## Install
Installs in ```$GOPATH/bin```
```
make install
```

## Run
```
go run cmd/gosdf-generate/main.go
go run cmd/gosdf-convert/main.go
```
after installation:
```
gosdf-convert
```